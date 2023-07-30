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
	port := configs.Configuration().RunPort
	fmt.Println("running on port: ", port)
	if err := srv.Run(port, v1.Router()); err != nil {
		logrus.Fatalf("error occured")
	}
}
