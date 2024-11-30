package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/cutlery47/employee-service/internal/config"
	"github.com/cutlery47/employee-service/internal/model"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(conf config.Postgres) *Repository {
	url := fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DB,
	)

	db, err := sql.Open("pgx", url)
	if err != nil {
		log.Fatalf("error while open db: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error while pinging db: %v", err)
	}

	return &Repository{
		db: db,
	}
}

func (r *Repository) GetEmployee(ctx context.Context, id int) (model.GetEmployeeResponse, error) {
	getEmployeeQuery := `
	SELECT (id, part, name, family_name, middle_name, phone, city, office, position, date_of_birth, unit)
	FROM employees AS e
	WHERE
	e.id = $1
	`

	response := model.GetEmployeeResponse{}

	var unit string

	row := r.db.QueryRowContext(ctx, getEmployeeQuery, id)
	if err := row.Scan(
		&response.Id,
		&response.Part,
		&response.Name,
		&response.FamilyName,
		&response.MiddleName,
		&response.Phone,
		&response.Office,
		&response.Position,
		&response.DateOfBirth,
		&unit,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.GetEmployeeResponse{}, ErrUserNotFound
		}
		return model.GetEmployeeResponse{}, err
	}

	getTeammatesQuery := `
	SELECT (unit, position, name, middle_name, family_name, part)
	FROM employees AS e
	WHERE
	e.unit = $1
	`

	rows, err := r.db.QueryContext(ctx, getTeammatesQuery, unit)
	if err != nil {
		return model.GetEmployeeResponse{}, err
	}

	teammates := []model.BaseEmployee{}
	for rows.Next() {
		teammate := model.BaseEmployee{}
		rows.Scan(
			&teammate.Unit,
			&teammate.Position,
			&teammate.Name,
			&teammate.MiddleName,
			&teammate.FamilyName,
			&teammate.Part,
		)

		teammates = append(teammates, teammate)
	}

	response.Teammates = teammates
	return response, nil
}

func (r *Repository) GetBaseEmployees(ctx context.Context, id int) (model.GetBaseEmployeesResponse, error) {
	// query := `
	// SELECT * FROM
	// employees AS e
	// WHERE e.id = $1
	// `

	// res := model.EmployeeMeta{}

	// row := r.db.QueryRowContext(ctx, query, id)
	// err := row.Scan(&res.Id, &res.Name, &res.Surname, &res.Department, &res.Role)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		return model.EmployeeMeta{}, ErrUserNotFound
	// 	}
	// 	return model.EmployeeMeta{}, err
	// }

	return model.GetBaseEmployeesResponse{}, nil
}

func (r *Repository) applyFilters(query string, filter model.GetBaseEmployeesFilter, limit, offset int, applied *[]any) (string, error) {
	// // контейнер со значениями структуры
	// st := reflect.ValueOf(filter)
	// // количество полей внутри структуры
	// numFields := st.NumField()
	// // количество примененных фильтров
	// filterCount := 0

	// // итерируемся по полям структуры
	// for i := 0; i < numFields; i++ {
	// 	if filterCount > 0 {
	// 		query += "AND\n"
	// 	}
	// 	filterCount++

	// 	structField :=

	// 	// достаем имя и значения поля структуры
	// 	fieldName := st.Field(i).Type().Name()
	// 	fieldValue := st.FieldByName(fieldName)

	// 	if fieldName == "DateOfBirth" {
	// 		// достаем из поля строку и конвертируем в time.Time
	// 		strTime := fieldValue.String()

	// 		parsedTime, err := time.Parse(time.DateOnly, strTime)
	// 		if err != nil {
	// 			return "", ErrWrongDateFormat
	// 		}

	// 		*applied = append(*applied, parsedTime)
	// 	} else {
	// 		*applied = append(*applied, fieldValue.String())
	// 	}

	// 	fieldTag := utils.GetStructTag()

	// 	query += fmt.Sprintf("%v = $%v", xyu, filterCount)
	// 	*applied = append(*applied, pizda)
	// }
	return "", nil
}
