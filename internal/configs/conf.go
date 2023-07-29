package configs

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Config struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	User     string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	DbName   string `env:"DBNAME"`
	SSlMode  string `env:"MODELESS"`
}

var Instance Config

func Configuration() *Config {
	// Load environment variables from the first file (.env)

	Instance = Config{
		Host:     "localhost",
		Port:     "5436",
		User:     "postgres",
		Password: "qwerty",
		DbName:   "postgres",
		SSlMode:  "disable",
	}
	fmt.Println("Instance")
	fmt.Println(Instance)
	return &Instance
}

func InitConfig() (string, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("configs")
	port := viper.GetString("port")
	fmt.Println(port)
	return port, viper.ReadInConfig()
}
func NewPostgresConfig(cfg *Config) (*sqlx.DB, error) {
	//postgres:qwerty@localhost:5436/postgres?sslmode=disable
	//cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SSlMode
	db, err := sqlx.Connect("postgres", "postgres:qwerty@localhost:5436/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if db != nil {
		panic("no database connection")
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
