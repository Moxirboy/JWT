package v1

import (
	"JWT/internal/configs"
	"JWT/internal/conn"
	"JWT/internal/controller/http/v1/handler"
	repository "JWT/internal/service/repo"
	"JWT/internal/service/usecase"
	json "JWT/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Router() *gin.Engine {
	db := repository.NewRepo(conn.Db)
	_ = json.GetterSigningKey()
	logrus.SetFormatter(new(logrus.JSONFormatter))
	server := usecase.NewService(db)
	handle := handler.NewHandler(server)
	if _, err := configs.InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	return handle.InitRoutes()
}
