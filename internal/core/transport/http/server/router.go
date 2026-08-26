package server

import (
	"net/http"

	"github.com/Vlad-6894/Activity_tracking/internal/core/transport/http/middleware"
)

type APIVersion string

const APIVersion1 APIVersion = "v1"

type PatternHandler struct {
	Pattern string
	Handler http.Handler
}

type APIVersionRouter struct {
	apiVersion APIVersion
	routes     []Route
	middleware []middleware.Middleware
}

func NewAPIVersionRouter(apiVersion APIVersion, mw ...middleware.Middleware) *APIVersionRouter {
	return &APIVersionRouter{
		apiVersion: apiVersion,
		middleware: mw,
	}
}

func (r *APIVersionRouter) AddRoutes(routes ...Route) {
	r.routes = append(r.routes, routes...)
}

func (r *APIVersionRouter) Handlers() []PatternHandler {
	handlers := make([]PatternHandler, 0, len(r.routes))

	for _, route := range r.routes {
		handlers = append(handlers, PatternHandler{
			Pattern: route.Method + " /api/" + string(r.apiVersion) + route.Path,
			Handler: middleware.Chain(route.WithMiddleware(), r.middleware...),
		})
	}

	return handlers
}
