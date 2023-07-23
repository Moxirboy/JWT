package repository

import (
	j "JWT"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user j.User) (int, error)
	GetUser(username string, password string) (j.User, error)
}
type Todolist interface {
}
type TodoItem interface {
}
type Repository struct {
	Authorization
	Todolist
	TodoItem
}

func NewRepo(db *sqlx.DB) *Repository {

	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
