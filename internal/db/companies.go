package db

import (
	"companies-test-task/internal/models"
	"context"
)

func (p *psql) CreateCompany(comp *models.Company) error {
	const query = `
	INSERT INTO companies (name, description, employees, registered, company_type)   
	VALUES ($1, $2, $3, $4, $5)
	RETURNING TEXT(id);`

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
	const query = `SELECT * FROM companies WHERE text(id)=$1;`

	var comp models.Company
	err := p.db.QueryRow(context.Background(), query, id).Scan(
		&comp.Id, &comp.Name, &comp.Description, &comp.EmployeesCount, &comp.Registered, &comp.CompanyType)
	if err != nil {
		return nil, err
	}

	return &comp, nil
}

func (p *psql) EditCompany(comp *models.Company) error {
	const query = `
	UPDATE companies
	SET name = $2,
	    description = $3,
	    employees = $4,
	    registered = $5,
	    company_type = $6
	WHERE TEXT(id)=$1;`

	_, err := p.db.Exec(context.Background(), query,
		comp.Id, comp.Name, comp.Description, comp.EmployeesCount, comp.Registered, comp.CompanyType)
	if err != nil {
		return err
	}

	return nil
}

func (p *psql) DeleteCompany(id string) error {
	const query = "DELETE FROM companies WHERE TEXT(id)=$1;"
	_, err := p.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}
