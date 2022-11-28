package service

import (
	"companies-test-task/internal/db"
	"companies-test-task/pkg/dto"
	"fmt"
)

// Companies interface exposes all functionality of Companies service
type Companies interface {
	CreateCompany(company *dto.Company) (*dto.Company, error)
	GetCompany(id string) (*dto.Company, error)
	EditCompany() error // TODO
	DeleteCompany(id string) error
}

// companiesSvc manages the internal state of Companies service
type companiesSvc struct {
	cfg     Config
	storage db.Storage
}

// New initializes Companies service handler
func New(cfg Config, storage db.Storage) (*companiesSvc, error) {
	svc := new(companiesSvc)

	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("invalid companies service configuration: %w", err)
	}
	svc.cfg = cfg
	svc.storage = storage

	return svc, nil
}
