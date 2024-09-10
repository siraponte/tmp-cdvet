// Server for the APIs.
package api

import (
	"cdvet/app/config"
	"cdvet/app/pkg/logger"
	"cdvet/app/pkg/router"
	"os"
)

func Run(cfg *config.Config, log logger.Logger) {
	r := router.NewRouter("/api")

	// Create openapi routes if enabled.
	if cfg.App.OpenAPI.Enabled {
		r.GroupPrefix("/docs", NewOpenapiRouter(cfg, log))
	}

	// Create versioned routes.
	r.GroupPrefix("/v1", NewRouterV1(cfg, log))

	// Start the server.
	log.Debugf("Backend available at %s", cfg.App.Http.BackendURL)
	log.Infof("Starting the server on http://%s", cfg.App.Http.Addr)
	err := r.ListenAndServe(cfg.App.Http.Addr)
	if err != nil {
		log.Errorf("Cannot start the server: %v", err)
		os.Exit(1)
	}
}
