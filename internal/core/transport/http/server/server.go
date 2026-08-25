package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Vlad-6894/Activity_tracking/internal/core/logger"
	"github.com/Vlad-6894/Activity_tracking/internal/core/transport/http/middleware"
	"go.uber.org/zap"
)

type Server struct {
	mux *http.ServeMux
	cfg Config
	log *logger.Logger

	mw []middleware.Middleware
}

func New(cfg Config, log *logger.Logger, mw ...middleware.Middleware) *Server {
	return &Server{
		mux: http.NewServeMux(),
		cfg: cfg,
		log: log,
		mw:  mw,
	}
}

func (s *Server) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		for _, handler := range router.Handlers() {
			s.mux.Handle(handler.Pattern, handler.Handler)
		}
	}
}

func (s *Server) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		s.mux.Handle(route.Method+" "+route.Path, route.WithMiddleware())
	}
}

func (s *Server) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:    s.cfg.Addr,
		Handler: middleware.Chain(s.mux, s.mw...),
	}

	errCh := make(chan error, 1)

	go func() {
		defer close(errCh)

		s.log.Info("start HTTP server", zap.String("addr", s.cfg.Addr))

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		return fmt.Errorf("http server %w", err)
	case <-ctx.Done():
		s.log.Info("http server shutting down")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			_ = srv.Close()
			return fmt.Errorf("http server shutdown %w", err)
		}

		s.log.Info("http server stopped")
		return nil
	}
}
