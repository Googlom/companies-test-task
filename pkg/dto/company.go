package dto

type Company struct {
	Id             string  `json:"id"`
	Name           string  `json:"name" binding:"required,gte=1,lte=15"`
	Description    *string `json:"description,omitempty" binding:"omitempty,lte=3000"`
	EmployeesCount int     `json:"employees_count"`
	Registered     bool    `json:"registered"`
	CompanyType    string  `json:"company_type"`
}
