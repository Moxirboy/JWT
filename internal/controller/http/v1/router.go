package v1

import (
	"JWT/internal/conn"
	"JWT/internal/controller/http/v1/handler"
	repository "JWT/internal/service/repo"
	"JWT/internal/service/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Router() *gin.Engine {
	db := repository.NewRepo(conn.Db)

	logrus.SetFormatter(new(logrus.JSONFormatter))
	server := usecase.NewService(db)
	handle := handler.NewHandler(server)
	defer func() {
		if err := recover(); err != nil {
			// Handle the panic here
			fmt.Println("Recovered from panic:", err)
		}
	}()
	return handle.InitRoutes()
}
