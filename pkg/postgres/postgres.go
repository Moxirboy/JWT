package postgres

import (
	"JWT/internal/configs"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewPostgresConfig(cfg *configs.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)

	instance, err := sqlx.Open("pgx", psqlString)
	if err != nil {
		return nil, err
	}
	return instance, nil
}
