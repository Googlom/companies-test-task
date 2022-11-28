package service

import (
	"companies-test-task/internal/models"
	"companies-test-task/pkg/dto"
)

func (c *companiesSvc) CreateCompany(company *dto.Company) (*dto.Company, error) {
	dbComp := models.CompanyFromDto(company)
	err := c.storage.CreateCompany(dbComp)
	if err != nil {
		return nil, err
	}

	return dbComp.ToDto(), nil
}

func (c *companiesSvc) GetCompany(id string) (*dto.Company, error) {
	dbCompany, err := c.storage.GetCompany(id)
	if err != nil {
		return nil, err
	}

	return dbCompany.ToDto(), nil
}

func (c *companiesSvc) EditCompany() error {
	//TODO implement me
	err := c.storage.EditCompany(nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *companiesSvc) DeleteCompany(id string) error {
	err := c.storage.DeleteCompany(id)
	if err != nil {
		return err
	}

	return nil
}
