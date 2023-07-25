package handler

import "JWT/internal/service/usecase"

type Handler struct {
	services *usecase.Service
}

func NewHandler(services *usecase.Service) *Handler {
	return &Handler{services: services}
}
