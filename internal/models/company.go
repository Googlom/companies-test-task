package models

import (
	"companies-test-task/internal/httpserver/errors"
	"companies-test-task/pkg/dto"
	"database/sql"
)

type Company struct {
	Id             string         `db:"id"`
	Name           string         `db:"name"`
	Description    sql.NullString `db:"description"`
	EmployeesCount int            `db:"employees"`
	Registered     bool           `db:"registered"`
	CompanyType    int            `db:"company_type"`
}

func (c *Company) ToDto() *dto.Company {
	return &dto.Company{
		Id:             c.Id,
		Name:           c.Name,
		Description:    nullStringToJson(c.Description),
		EmployeesCount: c.EmployeesCount,
		Registered:     c.Registered,
		CompanyType:    companyTypeToString(c.CompanyType),
	}
}

func (c *Company) Validate() error {
	if len(c.Name) < 2 || len(c.Name) > 15 {
		return &errors.APIError{Code: 400, Message: "name length must be in range from 2 to 15"}
	}
	if len(c.Description.String) > 3000 {
		return &errors.APIError{Code: 400, Message: "description length must not exceed 3000 chars"}
	}
	if c.CompanyType == UndefinedType {
		return &errors.APIError{Code: 400, Message: "invalid company type"}
	}

	return nil
}

func CompanyFromDto(dc *dto.Company) *Company {
	return &Company{
		Id:             dc.Id,
		Name:           dc.Name,
		Description:    jsonStringToNullString(dc.Description),
		EmployeesCount: dc.EmployeesCount,
		Registered:     dc.Registered,
		CompanyType:    companyTypeFromDto(dc.CompanyType),
	}
}

func nullStringToJson(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}

	return nil
}

func jsonStringToNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{
			String: *s,
			Valid:  true,
		}
	}

	return sql.NullString{}
}

func companyTypeToString(t int) string {
	return companyTypeEnum[t]
}

func companyTypeFromDto(s string) int {
	return companyTypeReverseEnum[s]
}

var companyTypeEnum = map[int]string{
	UndefinedType:          "undefined",
	CorporationType:        "Corporation",
	NonProfitType:          "NonProfit",
	CooperativeType:        "Cooperative",
	SoleProprietorshipType: "Sole Proprietorship",
}

var companyTypeReverseEnum = map[string]int{
	"undefined":           UndefinedType,
	"Corporation":         CorporationType,
	"NonProfit":           NonProfitType,
	"Cooperative":         CooperativeType,
	"Sole Proprietorship": SoleProprietorshipType,
}

const (
	UndefinedType = iota
	CorporationType
	NonProfitType
	CooperativeType
	SoleProprietorshipType
)
