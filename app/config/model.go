package config

import "encoding/json"

// Config is the main configuration struct.
type Config struct {
	Env string    `yaml:"env"`
	App AppConfig `yaml:"app"`
}

// PrettyPrint returns a pretty-printed JSON representation of the Config struct.
func (c *Config) PrettyPrint() string {
	s, _ := json.MarshalIndent(c, "", "\t")
	return string(s)
}

// AppConfig is the application configuration struct.
type AppConfig struct {
	OpenAPI OpenAPIConfig `yaml:"openapi"`
	Http    HttpConfig    `yaml:"http"`
	Logging LoggingConfig `yaml:"logging"`
}

// HttpConfig is the HTTP configuration struct.
type HttpConfig struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	BackendURL string `yaml:"backend_url"`
	Addr       string
	FullAddr   string
}

// OpenAPIConfig is the OpenAPI configuration struct.
type OpenAPIConfig struct {
	Enabled    bool   `yaml:"enabled"`
	DocFileLoc string `yaml:"doc_file_loc"`
}

// LoggingConfig is the logging configuration struct.
type LoggingConfig struct {
	Level string `yaml:"level"`
}
