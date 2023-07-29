package repository

import (
	j "JWT/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

type Authorization interface {
	CreateUsers(user j.User) (int, error)
}
type Repository struct {
	Authorization
}

func NewRepo(db *sqlx.DB) *Repository {

	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
