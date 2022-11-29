package service

import (
	"companies-test-task/internal/db"
	"companies-test-task/pkg/dto"
)

// Companies interface exposes all functionality of Companies service
type Companies interface {
	CreateCompany(company *dto.Company) (*dto.Company, error)
	GetCompany(id string) (*dto.Company, error)
	EditCompany(id string, patchJsonReq []byte) error
	DeleteCompany(id string) error
}

// Config holds the configuration of Companies service handler
type Config struct{}

// companiesSvc manages the internal state of Companies service
type companiesSvc struct {
	cfg     Config
	storage db.Storage
}

// New initializes Companies service handler
func New(cfg Config, storage db.Storage) *companiesSvc {
	svc := new(companiesSvc)
	svc.cfg = cfg
	svc.storage = storage
	return svc
}
