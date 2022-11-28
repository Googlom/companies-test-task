package db

import (
	"companies-test-task/internal/models"
	"context"
	"fmt"
)

func (p *psql) CreateCompany(comp *models.Company) error {
	const query = `
	INSERT INTO companies (name, description, employees, registered, company_type)   
	VALUES ($1, $2, $3, $4, $5)
	RETURNING TEXT(id)`

	id := ""
	err := p.db.QueryRow(context.Background(), query,
		comp.Name, comp.Description, comp.EmployeesCount, comp.Registered, comp.CompanyType).Scan(&id)
	if err != nil {
		return err
	}

	comp.Id = id
	return nil
}

func (p *psql) GetCompany(id string) (*models.Company, error) {
	const query = "SELECT * FROM companies WHERE text(id)=$1"

	var comp models.Company
	err := p.db.QueryRow(context.Background(), query, id).Scan(
		&comp.Id, &comp.Name, &comp.Description, &comp.EmployeesCount, &comp.Registered, &comp.CompanyType)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return &comp, nil
}

func (p *psql) EditCompany(company *models.Company) error {
	//TODO implement me
	return nil
}

func (p *psql) DeleteCompany(id string) error {
	const query = "DELETE FROM companies WHERE TEXT(id)=$1"
	_, err := p.db.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}

	return nil
}
