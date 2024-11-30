package model

// api/v1/employees (POST)

type GetBaseEmployeesRequest struct {
	Id       int    `json:"id,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Unit     string `json:"unit,omitempty"`
	Project  string `json:"project,omitempty"`
	Role     string `json:"role,omitempty"`
	Position string `json:"position,omitempty"`
	City     string `json:"city,omitempty"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

type GetBaseEmployeesResponse []BaseEmployee

type BaseEmployee struct {
	Id         int    `json:"id"`
	IsGeneral  bool   `json:"is_general"`
	Role       string `json:"role"`
	Name       string `json:"name"`
	FamilyName string `json:"family_name"`
	MiddleName string `json:"middle_name"`
	Position   string `json:"position"`
	Unit       string `json:"unit"`
}

// api/v1/employee (POST)

type GetEmployeeRequest struct {
	Id int `json:"id,omitempty"`
}

type GetEmployeeResponse struct {
	Id          int    `json:"id"`
	Role        string `json:"role"`
	Name        string `json:"name"`
	FamilyName  string `json:"family_name"`
	MiddleName  string `json:"middle_name"`
	Phone       string `json:"phone"`
	City        string `json:"city"`
	Project     string `json:"project"`
	Office      string `json:"office_address"`
	Position    string `json:"position"`
	DateOfBirth string `json:"birth_date"`
	Teammates   []BaseEmployee
}

// ----------------------

// type EmployeeMeta struct {
// 	EmployeeGeneral
// 	DateOfBirth time.Time
// 	Cellphone   string
// 	Email       string
// 	Address     string
// 	City        string
// 	Team        string
// 	Position    string
// }

// type EmployeeGeneral struct {
// 	Id         int
// 	Name       string
// 	Surname    string
// 	Department string
// 	Role       string
// }

type GetEmployeeCitiesResponse struct {
	Cities []string
}

type GetEmployeePositionsResponse struct {
	Positions []string
}

type GetEmployeeProjectsResponse struct {
	Projects []string
}

type GetEmployeeRolesResponse struct {
	Roles []string
}

type GetEmployeeUnitsResponse struct {
	Units []string
}
