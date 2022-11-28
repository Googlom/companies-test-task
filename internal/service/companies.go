package service

import (
	"companies-test-task/pkg/models"
)

func (c *companiesSvc) CreateCompany(company *models.Company) error {
	//TODO implement me
	err := c.storage.CreateCompany(company)
	if err != nil {
		return err
	}

	return nil
}

func (c *companiesSvc) GetCompany(id string) (*models.Company, error) {
	//TODO implement me
	cmp, err := c.storage.GetCompany(id)
	if err != nil {
		return nil, err
	}

	return cmp, nil
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
	//TODO implement me
	err := c.storage.DeleteCompany(id)
	if err != nil {
		return err
	}

	return nil
}
