package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cutlery47/employee-service/internal/config"
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

func (r *Repository) GetMeta() {

}
