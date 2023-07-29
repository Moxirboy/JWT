package configs

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/spf13/viper"
)

type Config struct {
	AppName    string `env:"APP_NAME"`
	AppVersion string `env:"APP_VERSION"`
	Postgres
	JWT
}
type Postgres struct {
	Port     uint16 `env:"POSTGRES_PORT"`
	Host     string `env:"POSTGRES_HOST"`
	Password string `env:"POSTGRES_PASSWORD"`
	User     string `env:"POSTGRES_USER"`
	Database string `env:"POSTGRES_DB"`
	PoolMax  int    `env:"POSTGRES_POOL_MAX"`
}
type JWT struct {
	SigningKey string `env:"SIGNING"`
	Salt       string `env:"SALT"`
	TokenTTL   string `env:"TOKENTTL"`
}

var instance Config

func Configuration() *Config {

	if err := env.Parse(&instance); err != nil {
		panic(err)
	}

	return &instance

}

func InitConfig() (string, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("configs")
	port := viper.GetString("port")
	fmt.Println(port)
	return port, viper.ReadInConfig()
}
