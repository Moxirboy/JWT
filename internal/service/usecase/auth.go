package usecase

import (
	"JWT/internal/models"
	"JWT/internal/service/repo"
)

type AuthService struct {
	repo repository.Authorization
}
type Authorization interface {
	CreateUser(user models.User) (int, error)
}
type Service struct {
	Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
