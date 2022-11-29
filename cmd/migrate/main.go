package main

import (
	"companies-test-task/internal/config"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
)

func main() {
	cfg := config.LoadDatabase()
	m, err := migrate.New(
		fmt.Sprintf("file://%s", cfg.MigrationsPath),
		fmt.Sprintf("pgx://%s:%s@%s:%d/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName),
	)
	if err != nil {
		log.Fatalf("open postgres connection failed: %s", err)
	}

	err = m.Up()

	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("migrate failed: %s", err)
		}
	}

	log.Println("migration complete.")
	os.Exit(0)
}
