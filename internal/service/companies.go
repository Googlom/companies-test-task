package service

import (
	"companies-test-task/internal/models"
	"companies-test-task/pkg/dto"
	"encoding/json"
	"fmt"

	jsonpatch "gopkg.in/evanphx/json-patch.v5"
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

func (c *companiesSvc) EditCompany(id string, patchJsonReq []byte) error {
	dbComp, err := c.storage.GetCompany(id)
	if err != nil {
		return fmt.Errorf("cannot get company: %w", err)
	}
	dbCompBytes, err := json.Marshal(dbComp.ToDto())
	if err != nil {
		return fmt.Errorf("marshal company: %w", err)
	}

	patchedCompBytes, err := jsonpatch.MergePatch(dbCompBytes, patchJsonReq)
	if err != nil {
		return fmt.Errorf("merge patch: %w", err)
	}

	var patchedComp dto.Company
	if err = json.Unmarshal(patchedCompBytes, &patchedComp); err != nil {
		return fmt.Errorf("unmarshal patched: %w", err)
	}

	if err = c.storage.EditCompany(models.CompanyFromDto(&patchedComp)); err != nil {
		return fmt.Errorf("company update failed: %w", err)
	}

	return nil
}

func (c *companiesSvc) DeleteCompany(id string) error {
	return c.storage.DeleteCompany(id)
}
