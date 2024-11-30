package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/cutlery47/employee-service/internal/config"
	"github.com/cutlery47/employee-service/internal/model"

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

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetEmployee(ctx context.Context, id int) (model.GetEmployeeResponse, error) {
	getEmployeeQuery := `
	SELECT (id, role, name, family_name, middle_name, phone, city, project, office_address, position, birth_date, unit_id)
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
	SELECT (e.id, is_general, role, name, family_name, middle_name, position, u.name)
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
		rows.Scan(
			&teammate.Id,
			&teammate.IsGeneral,
			&teammate.Role,
			&teammate.FamilyName,
			&teammate.MiddleName,
			&teammate.Position,
			&teammate.Unit,
		)

		teammates = append(teammates, teammate)
	}

	response.Teammates = teammates
	return response, nil
}

func (r *Repository) GetBaseEmployees(ctx context.Context, request model.GetBaseEmployeesRequest) (model.GetBaseEmployeesResponse, error) {
	getEmployeesQuery := `
	SELECT (e.id, is_general, role, name, family_name, middle_name, position, u.name)
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
