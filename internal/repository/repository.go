package repository

import (
	"database/sql"
	"log"
	"os"
)

type Repository struct {
	db *sql.DB
}

func NewRepository() *Repository {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
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
