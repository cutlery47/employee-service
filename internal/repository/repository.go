package repository

import (
	"context"
	"database/sql"
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

func (r *Repository) Get(ctx context.Context, filter model.GetBaseEmployeesFilter, limit, offset int) (model.GetBaseEmployeesResponse, error) {
	// query := `
	// SELECT (id, name, surname, department, role)
	// FROM employees AS e
	// WHERE
	// `

	// var appliedFilters []interface{}

	// filteredQuery, err := r.applyFilters(query, filter, limit, offset, &appliedFilters)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (r *Repository) GetMeta(ctx context.Context, id int) (model.GetBaseEmployeesResponse, error) {
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
