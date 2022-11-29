package dto

type Company struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	Description    *string `json:"description,omitempty"`
	EmployeesCount int     `json:"employees_count"`
	Registered     bool    `json:"registered"`
	CompanyType    string  `json:"company_type"`
}
