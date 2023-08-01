package postgres

import (
	"JWT/internal/configs"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPostgresConfig(cfg *configs.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SslMode)

	instance, err := sqlx.Open("pgx", psqlString)
	if err != nil {
		panic(err)
	}
	err = instance.Ping()
	if err != nil {
		panic(err)
	}

	return instance, nil
}
