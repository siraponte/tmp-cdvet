package config

import (
	"cdvet/app/pkg/utils"
	"log"
	"strings"
)

// validate validates the configuration.
// Takes a pointer to the Config struct as input.
func validate(cfg *Config) {
	validateEnv(cfg)
	validateAppHttp(cfg)
	validateAppLogging(cfg)
}

// validateEnv validates the environment.
// Takes a pointer to the Config struct as input.
func validateEnv(cfg *Config) {
	// Ensure the env is lowercase.
	cfg.Env = strings.ToLower(cfg.Env)

	validEnvs := []string{"dev", "test", "prod"}

	if !utils.ContainsString(validEnvs, cfg.Env) {
		log.Fatalf("Invalid environment '%s'. Must be one of %v", cfg.Env, validEnvs)
	}
}

// validateAppHttp validates the HTTP configuration.
// Takes a pointer to the Config struct as input.
func validateAppHttp(cfg *Config) {
	if cfg.App.Http.Host == "" {
		log.Fatalf("Host must be provided")
	}

	if cfg.App.Http.Port == 0 {
		log.Fatalf("Port must be provided")
	}
}

// validateAppLogging validates the logging configuration.
// Takes a pointer to the Config struct as input.
func validateAppLogging(cfg *Config) {
	// Ensure the logging level is lowercase.
	cfg.App.Logging.Level = strings.ToLower(cfg.App.Logging.Level)

	validLevels := []string{"debug", "info", "warn", "error"}
	if !utils.ContainsString(validLevels, cfg.App.Logging.Level) {
		log.Fatalf("Invalid logging level '%s'. Must be one of %v", cfg.App.Logging.Level, validLevels)
	}
}
