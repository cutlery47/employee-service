package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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
	SELECT e.id, is_general, role, name, family_name, middle_name, position, u.name)
	FROM employees AS e
	JOIN units AS u
	ON e.unit_id = u.id
	WHERE
	`

	var appliedFilters []interface{}

	filteredQuery := r.applyBaseEmployeeFilters(getEmployeesQuery, request, &appliedFilters)
	rows, err := r.db.QueryContext(ctx, filteredQuery, appliedFilters...)
	if err != nil {
		return model.GetBaseEmployeesResponse{}, err
	}

	response := model.GetBaseEmployeesResponse{}

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
		)
		if err != nil {
			return model.GetBaseEmployeesResponse{}, err
		}

		response = append(response, employee)
	}

	return response, nil
}

func (r *Repository) applyBaseEmployeeFilters(query string, request model.GetBaseEmployeesRequest, applied *[]any) string {
	filterCount := 0

	if request.Id != 0 {
		filterCount++
		query += fmt.Sprintf("id = $%v\n", filterCount)
		*applied = append(*applied, request.Id)
	}

	if request.FullName != "" {
		if filterCount > 0 {
			query += "AND\n"
		}
		filterCount++
		query += fmt.Sprintf("full_name = $%v\n", filterCount)
		*applied = append(*applied, request.FullName)
	}

	if request.Unit != "" {
		if filterCount > 0 {
			query += "AND\n"
		}
		filterCount++
		query += fmt.Sprintf("unit = $%v\n", filterCount)
		*applied = append(*applied, request.Unit)
	}

	if request.Project != "" {
		if filterCount > 0 {
			query += "AND\n"
		}
		filterCount++
		query += fmt.Sprintf("project = $%v\n", filterCount)
		*applied = append(*applied, request.Project)
	}

	if request.Role != "" {
		if filterCount > 0 {
			query += "AND\n"
		}
		filterCount++
		query += fmt.Sprintf("role = $%v\n", filterCount)
		*applied = append(*applied, request.Role)
	}

	if request.Position != "" {
		if filterCount > 0 {
			query += "AND\n"
		}
		filterCount++
		query += fmt.Sprintf("position = $%v\n", filterCount)
		*applied = append(*applied, request.Position)
	}

	if request.City != "" {
		if filterCount > 0 {
			query += "AND\n"
		}
		filterCount++
		query += fmt.Sprintf("city = $%v\n", filterCount)
		*applied = append(*applied, request.City)
	}

	filterCount++
	query += fmt.Sprintf("LIMIT $%v OFFSET $%v;", filterCount, filterCount+1)
	*applied = append(*applied, request.Limit, request.Offset)

	return query
}

func (r *Repository) GetHints(ctx context.Context, field string, value string) (interface{}, error) {
	getHintsQuery := `
	SELECT ($1)
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
	}
	return nil, nil
}
