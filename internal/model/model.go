package model

// api/v1/employees (POST)

type GetBaseEmployeesRequest struct {
	City     string
	Position string
	Part     string
	Project  string
	Unit     string
	FullName string
}

type GetBaseEmployeesResponse []BaseEmployee

type BaseEmployee struct {
	IsGeneral  bool
	Unit       string
	Position   string
	MiddleName string
	FamilyName string
	Name       string
	Part       string
}

// api/v1/employee (POST)

type GetEmployeeResponse struct {
	Id          int
	Part        string
	Name        string
	FamilyName  string
	MiddleName  string
	Phone       string
	City        string
	Office      string
	Position    string
	DateOfBirth string
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
