package main

import (
	s "JWT"
	"JWT/pkg/handler"
	repository "JWT/pkg/repo"
	"JWT/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	db, err := repository.NewPostgresConfig(repository.Configuration())
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err)
	}
	repos := repository.NewRepo(db)
	services := service.NewService(repos)
	handlers := handler.NewHAndler(services)
	srv := new(s.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured")
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
