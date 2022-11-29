package main

import (
	"companies-test-task/internal/db"
	"companies-test-task/internal/httpserver"
	"companies-test-task/internal/service"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	httpListenEnv       = "LISTEN_ADDR"
	jwtHmacSecretEnv    = "HMAC_SECRET"
	dbHostEnv           = "DB_HOST"
	dbPortEnv           = "DB_PORT"
	dbUserEnv           = "DB_USER"
	dbPasswordEnv       = "DB_PASSWORD"
	dbNameEnv           = "DB_NAME"
	dbMigrationsPathEnv = "DB_MIGRATIONS"
)

type configContainer struct {
	HttpServer httpserver.Config
	Service    service.Config
	Database   db.Config
}

func loadConfiguration() (*configContainer, error) {
	cfg := configContainer{
		HttpServer: httpserver.Config{
			HmacSecret: envOrDefaultString(jwtHmacSecretEnv, ""),
			ListenAddr: envOrDefaultString(httpListenEnv, ":8080"),
		},
		Service: service.Config{},
		Database: db.Config{
			Host:           envOrDefaultString(dbHostEnv, "localhost"),
			Port:           envOrDefaultInt(dbPortEnv, 5432),
			User:           envOrDefaultString(dbUserEnv, "companies"),
			Password:       envOrDefaultString(dbPasswordEnv, ""),
			DbName:         envOrDefaultString(dbNameEnv, "companies"),
			MigrationsPath: envOrDefaultString(dbMigrationsPathEnv, "db_migration"),
		},
	}
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validateConfig(cfg configContainer) error {
	if len(cfg.HttpServer.HmacSecret) == 0 {
		return fmt.Errorf("JWT HMAC secret must be set with '%s' env var", jwtHmacSecretEnv)
	}
	if len(cfg.Database.Password) == 0 {
		return fmt.Errorf("database password must be set with '%s' env var", dbPasswordEnv)
	}

	return nil
}

// envOrDefaultString tries to load string value from env variable using envKey.
// If it's not set then it returns defaultValue
func envOrDefaultString(envKey, defaultValue string) string {
	s := os.Getenv(envKey)
	if s != "" {
		return s
	}

	return defaultValue
}

// envOrDefaultInt tries to load int value from env variable using envKey.
// If it's not set or non-numeric value then it returns defaultValue
func envOrDefaultInt(envKey string, defaultValue int) int {
	s := os.Getenv(envKey)
	if s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Printf("invalid number %s, using default %d", s, defaultValue)
			return defaultValue
		}

		return i
	}

	return defaultValue
}
