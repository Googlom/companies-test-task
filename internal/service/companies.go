package service

import (
	"companies-test-task/pkg/models"
	"fmt"
)

// Companies interface exposes all functionality of Companies service
type Companies interface {
	CreateCompany(company *models.Company) error
	GetCompany(id string) (*models.Company, error)
	EditCompany() error // TODO
	DeleteCompany(id string) error
}

// companiesSvc manages the internal state of Companies service
type companiesSvc struct {
	cfg Config
}

// New initializes Companies service handler
func New(cfg Config) (*companiesSvc, error) {
	svc := new(companiesSvc)

	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("invalid companies service configuration: %w", err)
	}
	svc.cfg = cfg

	return svc, nil
}

func (c *companiesSvc) CreateCompany(company *models.Company) error {
	//TODO implement me
	return nil
}

func (c *companiesSvc) GetCompany(id string) (*models.Company, error) {
	//TODO implement me
	return nil, nil
}

func (c *companiesSvc) EditCompany() error {
	//TODO implement me
	return nil
}

func (c *companiesSvc) DeleteCompany(id string) error {
	//TODO implement me
	return nil
}
