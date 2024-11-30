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

func (r *Repository) Get() {

}

func (r *Repository) GetMeta(ctx context.Context, id int) (model.UserMeta, error) {
	query := `
	SELECT * FROM 
	employees AS e
	WHERE e.id = $1
	`

	res := model.UserMeta{}

	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&res.Id, &res.Name, &res.Surname, &res.Department, &res.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UserMeta{}, ErrUserNotFound
		}
		return model.UserMeta{}, err
	}

	return model.UserMeta{}, nil
}
