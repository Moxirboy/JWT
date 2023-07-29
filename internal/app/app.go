package app

import (
	"JWT/internal/configs"
	v1 "JWT/internal/controller/http/v1"
	"JWT/internal/server"
	"fmt"
	"github.com/sirupsen/logrus"
)

func App() {
	srv := new(server.Server)
	port, err := configs.InitConfig()
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err)
	}
	fmt.Println("port")
	fmt.Println(port)
	if err := srv.Run("5432", v1.Router()); err != nil {
		logrus.Fatalf("error occured")
	}
}
