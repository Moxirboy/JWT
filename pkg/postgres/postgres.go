package postgres

import (
	"JWT/internal/configs"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewPostgresConfig(cfg *configs.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.PoolMax)

	instance, err := sqlx.Open("pgx", psqlString)
	if err != nil {
		return nil, err
	}
	return instance, nil
}
