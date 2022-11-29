package service

import (
	"companies-test-task/internal/db/mock"
	"companies-test-task/internal/models"
	"database/sql"
	"reflect"
	"testing"
)

func Test_companiesSvc_EditCompany(t *testing.T) {
	tests := []struct {
		name          string
		companyToEdit models.Company
		patchJsonReq  []byte
		wantComp      models.Company
		wantErr       bool
	}{
		{
			name: "normal edit",
			companyToEdit: models.Company{
				Name:           "TestName",
				Description:    sql.NullString{String: "Test descr", Valid: true},
				EmployeesCount: 1,
				Registered:     true,
				CompanyType:    1,
			},
			patchJsonReq: []byte(`{
			  "name": "NewTestName",
              "description": null,
			  "employees_count": 55
			}`),
			wantComp: models.Company{
				Name:           "NewTestName",
				Description:    sql.NullString{},
				EmployeesCount: 55,
				Registered:     true,
				CompanyType:    1,
			},
			wantErr: false,
		},
		{
			name: "invalid name",
			companyToEdit: models.Company{
				Name:           "TestName",
				EmployeesCount: 1,
				Registered:     true,
				CompanyType:    1,
			},
			patchJsonReq: []byte(`{
			  "name": "Toooooo Looooong Name"
			}`),
			wantComp: models.Company{
				Name:           "TestName",
				EmployeesCount: 1,
				Registered:     true,
				CompanyType:    1,
			},
			wantErr: true,
		},
		{
			name: "invalid company type",
			companyToEdit: models.Company{
				Name:           "TestName",
				EmployeesCount: 1,
				Registered:     true,
				CompanyType:    1,
			},
			patchJsonReq: []byte(`{
			  "company_type": "Not a company type"
			}`),
			wantComp: models.Company{
				Name:           "TestName",
				EmployeesCount: 1,
				Registered:     true,
				CompanyType:    1,
			},
			wantErr: true,
		},
		{
			name: "empty patch",
			companyToEdit: models.Company{
				Name:           "TestName",
				EmployeesCount: 1,
				Registered:     true,
				CompanyType:    1,
			},
			patchJsonReq: []byte(`{}`),
			wantComp: models.Company{
				Name:           "TestName",
				EmployeesCount: 1,
				Registered:     true,
				CompanyType:    1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &companiesSvc{
				cfg:     Config{},
				storage: mock.NewMockDb(),
			}
			_ = c.storage.CreateCompany(&tt.companyToEdit)

			if err := c.EditCompany(tt.companyToEdit.Id, tt.patchJsonReq); (err != nil) != tt.wantErr {
				t.Errorf("EditCompany() error = %v, wantErr %v", err, tt.wantErr)
			}

			resultComp, _ := c.storage.GetCompany(tt.companyToEdit.Id)
			resultComp.Id = "" // we don't care about Id value
			if !reflect.DeepEqual(*resultComp, tt.wantComp) {
				t.Errorf("EditCompany() result = %v, want %v", *resultComp, tt.wantComp)
			}
		})
	}
}
