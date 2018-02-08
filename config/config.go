package config

import (
	"os"
)

// Config represents configuration option state for the app.
type Config struct {
	Host             string
	Port             string
	Domain           string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPass     string
	PostgresDatabase string
}

// FromEnv creates a configuration from the environment.
func FromEnv() *Config {
	return &Config{
		Host:         os.Getenv("TESLA_HOST"),
		Port:         os.Getenv("TESLA_PORT"),
		PostgresHost: os.Getenv("TESLA_POSTGRESHOST"),
		PostgresPort: os.Getenv("TESLA_POSTGRESPORT"),
		PostgresUser: os.Getenv("TESLA_POSTGRESUSER"),
		PostgresPass: os.Getenv("TESLA_POSTGRESPASS"),
	}
}

// LocalTest creates a configuration using test values for local development.
func LocalTest() *Config {
	return &Config{
		Host:             "localhost",
		Port:             "4000",
		PostgresHost:     "localhost",
		PostgresPort:     "5432",
		PostgresUser:     "api",
		PostgresPass:     "",
		PostgresDatabase: "api",
	}
}
