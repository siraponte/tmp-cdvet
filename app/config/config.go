// Provides configuration struct from config file parsing and validation.
package config

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"gopkg.in/yaml.v3"
)

func Init() *Config {
	// Define the flags
	configPath := flag.String("config", "./config/application.yaml", "Path to the YAML configuration file")
	validate := flag.Bool("validate", false, "If set to true, validate the config file and exit")
	host := flag.String("host", "", "Force the host to be used; Overrides the host in the config file")
	port := flag.Int("port", 0, "Force the port to be used; Overrides the port in the config file")

	// Parse the flags
	flag.Parse()

	// Parse the config file
	cfg, err := parse(configPath)
	if err != nil {
		log.Fatalf("Invalid configuration file '%s': %v", *configPath, err)
	}

	// Override the host if the flag has been set
	if *host != "" {
		cfg.App.Http.Host = *host
	}

	// Override the port if the flag has been set
	if *port != 0 {
		cfg.App.Http.Port = *port
	}

	// If validating flag has been set to true, print the config and exit
	if *validate {
		println(cfg.PrettyPrint())
		println("Configuration file is valid")
		os.Exit(0)
	}

	return cfg
}

/* -------------------------------------------------------------------------- */
/*                                  INTERNAL                                  */
/* -------------------------------------------------------------------------- */

func parse(configFile *string) (*Config, error) {
	// Create a new Config struct
	cfg := &Config{}

	// Read config file
	data, err := os.ReadFile(*configFile)
	if err != nil {
		return nil, err
	}

	// Unmarshal the yaml config file into the cfg struct
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Cannot unmarshal configuration file: %v", err)
		return nil, err
	}

	// Validate the config
	validate(cfg)

	// Add composite fields
	cfg.App.Http.Addr = createAddr(cfg.App.Http.Host, cfg.App.Http.Port)
	cfg.App.Http.FullAddr = createFullAddr(cfg.App.Http.Addr)

	return cfg, nil
}

// createAddr creates an address from the protocol, host, and port.
// Takes the protocol, host, and port as input and returns the address as a string.
func createAddr(host string, port int) string {
	portStr := fmt.Sprintf("%d", port)

	return net.JoinHostPort(host, portStr)
}

func createFullAddr(addr string) string {
	return fmt.Sprintf("http://%s/api", addr)
}
