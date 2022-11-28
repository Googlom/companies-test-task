package db

import (
	"companies-test-task/internal/models"
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func (p *psql) CreateCompany(company *models.Company) error {
	//TODO implement me
	return nil
}

func (p *psql) GetCompany(id string) (*models.Company, error) {
	const query = "SELECT * FROM companies WHERE text(id)=$1"

	var comp models.Company
	err := pgxscan.Get(context.Background(), p.db, &comp, query, id)
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
	//TODO implement me
	return nil
}
