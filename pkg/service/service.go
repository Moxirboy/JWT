package service

import (
	jwt "JWT"
	repository "JWT/pkg/repo"
)

type Authorization interface {
	CreateUser(user jwt.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
