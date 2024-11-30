package model

type GetBaseEmployeesFilter struct {
	Name        *string `db:"name"`
	Surname     *string `db:"surname"`
	Department  *string `db:"department"`
	Role        *string `db:"role"`
	DateOfBirth *string `db:"date_of_birth"`
	Cellphone   *string `db:"cellphone"`
	Email       *string `db:"email"`
	Address     *string `db:"address"`
	City        *string `db:"city"`
	Team        *string `db:"team"`
	Position    *string `db:"position"`
}

func GetBaseEmployeesFromURL(values JSON) (GetBaseEmployeesFilter, error) {
	filter := GetBaseEmployeesFilter{}

	switch {
	case values.Has("name"):
		role := values.Get("name")
		filter.Name = &role
		fallthrough
	case values.Has("surname"):
		surname := values.Get("surname")
		filter.Surname = &surname
		fallthrough
	case values.Has("department"):
		department := values.Get("department")
		filter.Department = &department
		fallthrough
	case values.Has("role"):
		role := values.Get("role")
		filter.Role = &role
		fallthrough
	case values.Has("date_of_birth"):
		dateOfBirth := values.Get("date_of_birth")
		filter.DateOfBirth = &dateOfBirth
		fallthrough
	case values.Has("cellphone"):
		department := values.Get("department")
		filter.Cellphone = &department
		fallthrough
	case values.Has("email"):
		email := values.Get("email")
		filter.Email = &email
		fallthrough
	case values.Has("address"):
		address := values.Get("address")
		filter.Email = &address
		fallthrough
	case values.Has("city"):
		city := values.Get("city")
		filter.City = &city
		fallthrough
	case values.Has("team"):
		team := values.Get("team")
		filter.Team = &team
		fallthrough
	case values.Has("position"):
		position := values.Get("position")
		filter.Position = &position
	}

	return filter, nil
}
