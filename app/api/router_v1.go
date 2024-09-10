package api

import (
	"cdvet/app/api/handler"
	"cdvet/app/config"
	"cdvet/app/pkg/logger"
	"cdvet/app/pkg/router"
)

func NewRouterV1(cfg *config.Config, log logger.Logger) router.GroupMember {
	return func(r router.Router) {
		// Initialize middlewares.
		middlewares := []router.Middleware{
			router.LoggingMiddleware(log),
		}

		r.Use(middlewares...)

		// App routes.
		r.Get("/health", handler.Health)

		r.GroupPrefix("/users", func(r router.Router) {})
	}
}
