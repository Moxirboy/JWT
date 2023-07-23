package repository

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

const (
	usersTable = "users"
)

type Config struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	User     string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	DbName   string `env:"DBNAME"`
	SSlMode  string `env:"SSLMODE"`
}

var Instance Config

func Configuration() *Config {
	if err := env.Parse(&Instance); err != nil {
		panic(err)
	}
	return &Instance
}
func NewPostgresConfig(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SSlMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
