package mock

import (
	"companies-test-task/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type mockDb struct {
	companies map[string]models.Company
}

func NewMockDb() *mockDb {
	return &mockDb{
		companies: make(map[string]models.Company),
	}
}

func (m *mockDb) CreateCompany(comp *models.Company) error {
	if err := comp.Validate(); err != nil {
		return err
	}

	id := uuid.New().String()
	comp.Id = id
	m.companies[id] = *comp
	return nil
}

func (m *mockDb) GetCompany(id string) (*models.Company, error) {
	comp, found := m.companies[id]
	if !found {
		return nil, pgx.ErrNoRows
	}

	return &comp, nil
}

func (m *mockDb) EditCompany(comp *models.Company) error {
	if err := comp.Validate(); err != nil {
		return err
	}
	
	m.companies[comp.Id] = *comp
	return nil
}

func (m *mockDb) DeleteCompany(id string) error {
	delete(m.companies, id)
	return nil
}
