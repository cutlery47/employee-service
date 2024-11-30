package model

// api/v1/employees (POST)

type GetBaseEmployeesRequest struct {
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

// api/v1/hint (POST)

type GetHintRequest struct {
	City     string `json:"city_search_term,omitempty"`
	Position string `json:"position_search_term,omitempty"`
	Project  string `json:"project_search_term,omitempty"`
	Role     string `json:"role_search_term,omitempty"`
	Unit     string `json:"unit_search_term,omitempty"`
	Name     string `json:"name_search_term,omitempty"`
}

type GetEmployeeCitiesResponse struct {
	Cities []string `json:"cities"`
}

type GetEmployeePositionsResponse struct {
	Positions []string `json:"positions"`
}

type GetEmployeeProjectsResponse struct {
	Projects []string `json:"projects"`
}

type GetEmployeeRolesResponse struct {
	Roles []string `json:"roles"`
}

type GetEmployeeUnitsResponse struct {
	Units []string `json:"units"`
}

type GetEmployeeNamesResponse struct {
	Names []string `json:"names"`
}

type Unit struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// api/v1/unit (POST)

type GetUnitRequest struct {
	Id int `json:"id,omitempty"`
}

type GetUnitResponse struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	ParentId       int    `json:"parent_id"`
	LeaderFullName string `json:"leader_full_name"`
	Partisipants   []BaseEmployee
	Units          []Unit
}
