package main

import (
	"cdvet/app/api"
	"cdvet/app/config"
	"cdvet/app/pkg/logger"
)

// @title CDVet Swagger API
// @version 1.0
// @description Welcome to CDVet Documentation.
// @termsOfService http://swagger.io/terms/
//
// @contact.name Ko2
// @contact.url https://ko2.it
// @contact.email developers@ko2.it
//
// @BasePath /api
func main() {
	// Initialize the configuration.
	cfg := config.Init()

	// Initialize the logger.
	log := logger.New(&cfg.App.Logging)

	// Run the API server.
	api.Run(cfg, log)
}
