package db

import (
	"companies-test-task/internal/models"
	"context"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Storage exposes db methods for interacting with stored company items
type Storage interface {
	CreateCompany(company *models.Company) error
	GetCompany(id string) (*models.Company, error)
	EditCompany(company *models.Company) error
	DeleteCompany(id string) error
}

type Config struct {
	Host     string // host addr
	Port     int    // port
	User     string // user
	Password string // user password
	DbName   string // companies database name

	MigrationsPath string // directory where migration scripts are stored
}

// psql implements Postgres adapter for Companies storage
type psql struct {
	db *pgxpool.Pool
}

func New(cfg Config) (*psql, error) {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName,
	)
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("open postgres connection failed: %w", err)
	}

	if err = runMigrations(cfg); err != nil {
		return nil, fmt.Errorf("database migrations failed: %w", err)
	}

	return &psql{db: db}, nil
}

func runMigrations(cfg Config) error {
	m, err := migrate.New(
		fmt.Sprintf("file://%s", cfg.MigrationsPath),
		fmt.Sprintf("pgx://%s:%s@%s:%d/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName),
	)
	if err != nil {
		return fmt.Errorf("open postgres connection failed: %w", err)
	}

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	return err
}
