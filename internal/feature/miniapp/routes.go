package miniapp

import (
	"net/http"

	"github.com/Vlad-6894/Activity_tracking/internal/core/transport/http/server"
)

func (h *Handler) Routes() []server.Route {
	return []server.Route{
		{Method: http.MethodPost, Path: "/miniapp/me", Handler: h.me},
	}
}

func (h *Handler) RootRoutes() []server.Route {
	return []server.Route{
		{Method: http.MethodGet, Path: "/{$}", Handler: h.index}, // {$} означает строгий конец строки
		{Method: http.MethodGet, Path: "/ping", Handler: h.ping},
	}
}
