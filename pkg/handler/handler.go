package handler

import (
	s "JWT/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *s.Service
}

func NewHAndler(services *s.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	return router
}
