package repository

import (
	"JWT/internal/conn"
	_const "JWT/internal/const"
	j "JWT/internal/models"
	"fmt"
)

func (r *AuthPostgres) CreateUsers(user j.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name ,username,password_hash) values($1,$2,$3) RETURNING id", _const.UsersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func GetUser(username string, password string) (j.User, error) {
	var user j.User
	query := fmt.Sprintf("SELECT id 	FROM %s WHERE username = $1 AND password_hash =$2", _const.UsersTable)
	err := conn.Db.Get(&user, query, username, password)
	return user, err
}
