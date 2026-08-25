package server

import (
	"net/http"

	"github.com/Vlad-6894/Activity_tracking/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware []middleware.Middleware
}

func (r *Route) WithMiddleware() http.Handler {
	return middleware.Chain(r.Handler, r.Middleware...)
}
