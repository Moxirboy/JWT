package app

import (
	"JWT/internal/configs"
	v1 "JWT/internal/controller/http/v1"
	"JWT/internal/server"
	"github.com/sirupsen/logrus"
)

func App() {

	srv := new(server.Server)
	port, err := configs.InitConfig()
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err)
	}
	if err := srv.Run(port, v1.Router()); err != nil {
		logrus.Fatalf("error occured")
	}
}
