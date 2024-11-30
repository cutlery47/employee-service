package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/cutlery47/employee-service/internal/config"
	"github.com/cutlery47/employee-service/internal/model"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/sirupsen/logrus"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(conf config.Postgres) (*Repository, error) {
	url := fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DB,
	)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db.Ping: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("postgers.WithInstance: %v", err)
	}

	migrations := fmt.Sprintf("file://%v", conf.Migrations)
	m, err := migrate.NewWithDatabaseInstance(migrations, conf.DB, driver)
	if err != nil {
		return nil, fmt.Errorf("migrate.NewWithDatabaseInstance: %v", err)
	}

	logrus.Debug("applying migrations...")
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logrus.Debug("nothing to migrate")
		} else {
			return nil, fmt.Errorf("error when migrating: %v", err)
		}
	} else {
		logrus.Debug("migrated successfully!")
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetEmployee(ctx context.Context, id int) (model.GetEmployeeResponse, error) {
	getEmployeeQuery := `
	SELECT e.id, e.role_name, e.name, e.family_name, e.middle_name, e.phone, e.city, e.project, e.office_address, e.position, e.birth_date, e.unit_id
	FROM employees AS e
	WHERE
	e.id = $1
	`

	response := model.GetEmployeeResponse{}
	var unitId int

	row := r.db.QueryRowContext(ctx, getEmployeeQuery, id)
	if err := row.Scan(
		&response.Id,
		&response.Role,
		&response.Name,
		&response.FamilyName,
		&response.MiddleName,
		&response.Phone,
		&response.City,
		&response.Project,
		&response.Office,
		&response.Position,
		&response.DateOfBirth,
		&unitId,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.GetEmployeeResponse{}, ErrUserNotFound
		}
		return model.GetEmployeeResponse{}, err
	}

	getTeammatesQuery := `
	SELECT e.id, e.is_general, e.role_name, e.name, e.family_name, e.middle_name, e.position, u.name
	FROM employees AS e
	JOIN units AS u
	ON e.unit_id = u.id
	WHERE
	u.id = $1
	`

	rows, err := r.db.QueryContext(ctx, getTeammatesQuery, unitId)
	if err != nil {
		return model.GetEmployeeResponse{}, err
	}

	teammates := []model.BaseEmployee{}
	for rows.Next() {
		teammate := model.BaseEmployee{}
		if err := rows.Scan(
			&teammate.Id,
			&teammate.IsGeneral,
			&teammate.Role,
			&teammate.Name,
			&teammate.FamilyName,
			&teammate.MiddleName,
			&teammate.Position,
			&teammate.Unit,
		); err != nil {
			return model.GetEmployeeResponse{}, err
		}

		if teammate.Id != id {
			teammates = append(teammates, teammate)
		}
	}

	response.Teammates = teammates
	return response, nil
}

func (r *Repository) GetBaseEmployees(ctx context.Context, request model.GetBaseEmployeesRequest) (model.GetBaseEmployeesResponse, error) {
	getEmployeesQuery := `
	SELECT e.id, e.is_general, e.role_name, e.name, e.family_name, e.middle_name, e.position, u.name, ur.parent_id
	FROM employees AS e
	JOIN units AS u
	ON e.unit_id = u.id
	JOIN units_relations AS ur 
	ON u.id = ur.child_id
	WHERE
	`

	var appliedFilters []interface{}

	filteredQuery, err := r.applyBaseEmployeeFilters(getEmployeesQuery, request, &appliedFilters)
	if err != nil {
		return model.GetBaseEmployeesResponse{}, err
	}

	rows, err := r.db.QueryContext(ctx, filteredQuery, appliedFilters...)
	if err != nil {
		return model.GetBaseEmployeesResponse{}, err
	}

	var parent_id int
	employees := model.GetBaseEmployeesResponse{}

	for rows.Next() {
		employee := model.BaseEmployee{}
		err := rows.Scan(
			&employee.Id,
			&employee.IsGeneral,
			&employee.Role,
			&employee.Name,
			&employee.FamilyName,
			&employee.MiddleName,
			&employee.Position,
			&employee.Unit,
			&parent_id,
		)
		if err != nil {
			return model.GetBaseEmployeesResponse{}, err
		}

		employees = append(employees, employee)
	}

	if len(employees) == 0 {
		return model.GetBaseEmployeesResponse{}, ErrUserNotFound
	}

	getParentName := `
	SELECT u.name 
	FROM units AS u
	JOIN units_relations AS ur
	ON u.id = ur.child_id
	WHERE 
	ur.child_id = $1
	`

	for i := 0; i < len(employees); i++ {
		var parentUnit string
		parentMap := make(map[string]string)

		row, err := r.db.QueryContext(ctx, getParentName, parent_id)
		if err != nil {
			return model.GetBaseEmployeesResponse{}, err
		}

		if err := row.Scan(&parentUnit); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				parentMap[employees[i].Unit] = ""
			} else {
				return model.GetBaseEmployeesResponse{}, nil
			}
		}
		parentMap[employees[i].Name] = parentUnit
		employees[i].StringDict = parentMap
	}

	return employees, nil
}

func (r *Repository) applyBaseEmployeeFilters(query string, request model.GetBaseEmployeesRequest, applied *[]any) (string, error) {
	filterCount := 0

	if request.Unit != "" {
		if filterCount > 0 {
			query += "AND\n"
		}
		filterCount++
		query += fmt.Sprintf("u.name = $%v\n", filterCount)
		*applied = append(*applied, request.Unit)
	}

	if request.Project != "" {
		if filterCount > 0 {
			query += "AND\n"
		}
		filterCount++
		query += fmt.Sprintf("e.project = $%v\n", filterCount)
		*applied = append(*applied, request.Project)
	}

	if request.Role != "" {
		if filterCount > 0 {
			query += "AND\n"
		}
		filterCount++
		query += fmt.Sprintf("e.role_name = $%v\n", filterCount)
		*applied = append(*applied, request.Role)
	}

	if request.Position != "" {
		if filterCount > 0 {
			query += "AND\n"
		}
		filterCount++
		query += fmt.Sprintf("e.position = $%v\n", filterCount)
		*applied = append(*applied, request.Position)
	}

	if request.City != "" {
		if filterCount > 0 {
			query += "AND\n"
		}
		filterCount++
		query += fmt.Sprintf("e.city = $%v\n", filterCount)
		*applied = append(*applied, request.City)
	}

	if request.FullName != "" {
		if filterCount > 0 {
			query += "AND\n"
		}
		names := strings.Split(request.FullName, " ")

		if len(names) == 3 {
			name_1, name_2, name_3 := names[0], names[1], names[2]

			filterCount += 3
			query += fmt.Sprintf(
				"e.name IN ($%v, $%v, $%v) OR e.family_name IN ($%v, $%v, $%v) OR e.middle_name IN ($%v, $%v, $%v)\n",
				filterCount-2,
				filterCount-1,
				filterCount,
				filterCount-2,
				filterCount-1,
				filterCount,
				filterCount-2,
				filterCount-1,
				filterCount,
			)
			*applied = append(*applied, name_1, name_2, name_3)
		} else if len(names) == 2 {
			name_1, name_2 := names[0], names[1]

			filterCount += 2
			query += fmt.Sprintf(
				"e.name IN ($%v, $%v) OR e.family_name IN ($%v, $%v) OR e.middle_name IN ($%v, $%v)\n",
				filterCount-1,
				filterCount,
				filterCount-1,
				filterCount,
				filterCount-1,
				filterCount,
			)
			*applied = append(*applied, name_1, name_2)
		} else if len(names) == 1 {
			name := names[0]

			filterCount++
			query += fmt.Sprintf(
				"e.name = $%v OR e.family_name = $%v OR e.middle_name = $%v\n",
				filterCount,
				filterCount,
				filterCount,
			)
			*applied = append(*applied, name)
		} else {
			return "", ErrNameLengthExceeded
		}

	}

	if filterCount == 0 {
		query = strings.TrimSuffix(query, "WHERE\n\t")
	}

	filterCount++
	query += fmt.Sprintf("LIMIT $%v OFFSET $%v;", filterCount, filterCount+1)
	*applied = append(*applied, request.Limit, request.Offset)

	return query, nil
}

func (r *Repository) GetHints(ctx context.Context, field string, value string) (interface{}, error) {
	getHintsQuery := `
	SELECT $1
	FROM employees AS e
	WHERE
	e.$1 LIKE '$2%';
	`

	rows, err := r.db.QueryContext(ctx, getHintsQuery, field, value)

	if err != nil {
		return nil, err
	}

	var hints []string
	for rows.Next() {
		var hint string
		rows.Scan(hint)
		hints = append(hints, hint)
	}
	if field == "city" {
		var citiesResponse model.GetEmployeeCitiesResponse
		citiesResponse.Cities = hints
		return citiesResponse, nil
	} else if field == "position" {
		var positionsResponse model.GetEmployeePositionsResponse
		positionsResponse.Positions = hints
		return positionsResponse, nil
	} else if field == "project" {
		var projectsResponse model.GetEmployeeProjectsResponse
		projectsResponse.Projects = hints
		return projectsResponse, nil
	} else if field == "role" {
		var rolesResponse model.GetEmployeeRolesResponse
		rolesResponse.Roles = hints
		return rolesResponse, nil
	} else if field == "unit" {
		var unitsResponse model.GetEmployeeUnitsResponse
		unitsResponse.Units = hints
		return unitsResponse, nil
	} else if field == "name" {
		var namesResponse model.GetEmployeeNamesResponse
		namesResponse.Names = hints
		return namesResponse, nil
	}
	return nil, nil
}

func (r *Repository) GetUnit(ctx context.Context, id int) (model.GetUnitResponse, error) {
	getLeaderNameQuery := `
	SELECT name, family_name, middle_name
	FROM employees AS e
	WHERE
	e.unit_id = $1 AND e.is_general = TRUE;
	`
	getUnitNameQuery := `
	SELECT name
	FROM units AS u
	WHERE
	u.id = $1;
	`
	getParentIdQuery := `
	SELECT parent_id
	FROM unit_relation AS ur
	WHERE
	ur.child_id = $1;
	`
	getParticipantsQuery := `
	SELECT id, is_general, role, name, family_name, middle_name, position
	FROM employees AS e
	WHERE
	e.unit_id = $1;
	`
	getUnitsQuery := `
	SELECT u.id, u.name
	FROM units as u
	JOIN unit_relations as ur
	ON u.id = ur.child_id
	WHERE ur.parent_id = $1
	`

	var (
		response                                       model.GetUnitResponse
		leaderName, leaderFamilyName, leaderMiddleName string
		unitName                                       string
		parentId                                       int
		participants                                   []model.BaseEmployee
		units                                          []model.Unit
	)

	row := r.db.QueryRowContext(ctx, getLeaderNameQuery, id)
	if err := row.Scan(leaderName, leaderFamilyName, leaderMiddleName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.GetUnitResponse{}, ErrUserNotFound
		}
		return model.GetUnitResponse{}, err
	}
	row = r.db.QueryRowContext(ctx, getUnitNameQuery, id)
	if err := row.Scan(unitName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.GetUnitResponse{}, ErrUserNotFound
		}
		return model.GetUnitResponse{}, err
	}
	row = r.db.QueryRowContext(ctx, getParentIdQuery, id)
	if err := row.Scan(parentId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.GetUnitResponse{}, ErrUserNotFound
		}
		return model.GetUnitResponse{}, err
	}

	rows, err := r.db.QueryContext(ctx, getParticipantsQuery, id)
	if err != nil {
		return model.GetUnitResponse{}, err
	}

	for rows.Next() {
		participant := model.BaseEmployee{}
		rows.Scan(
			&participant.Id,
			&participant.IsGeneral,
			&participant.Role,
			&participant.FamilyName,
			&participant.MiddleName,
			&participant.Position,
		)
		participant.Unit = unitName
		participants = append(participants, participant)
	}

	rows, err = r.db.QueryContext(ctx, getUnitsQuery, id)
	if err != nil {
		return model.GetUnitResponse{}, err
	}

	for rows.Next() {
		unit := model.Unit{}
		rows.Scan(
			&unit.Id,
			&unit.Name,
		)
		units = append(units, unit)
	}

	response.Id = id
	response.Name = unitName
	response.ParentId = parentId
	response.LeaderFullName = fmt.Sprintf("%s %s %s", leaderName, leaderFamilyName, leaderMiddleName)
	response.Partisipants = participants
	response.Units = units

	return response, nil
}
