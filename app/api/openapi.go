package api

import (
	"cdvet/app/config"
	"cdvet/app/pkg/logger"
	"cdvet/app/pkg/router"
	"net/http"
	"net/url"

	_ "cdvet/openapi" // Required for swagger docs.

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewOpenapiRouter(cfg *config.Config, log logger.Logger) router.GroupMember {
	return func(r router.Router) {
		swaggerURL, err := url.JoinPath(cfg.App.Http.BackendURL, "/docs/swagger.json")
		if err != nil {
			log.Warnf("Documentation not available: Cannot join swagger URL: %v", err)
			return
		}
		indexURL, err := url.JoinPath(cfg.App.Http.BackendURL, "/docs/index.html")
		if err != nil {
			log.Warnf("Documentation not available: Cannot join index URL: %v", err)
			return
		}

		log.Debugf("OpenAPI documentation available at %s", indexURL)

		r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, cfg.App.OpenAPI.DocFileLoc)
		})

		r.Get("/{steps...}", httpSwagger.Handler(
			httpSwagger.URL(swaggerURL),
		))
	}
}
