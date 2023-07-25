package usecase

import (
	"JWT/internal/models"

	json "JWT/pkg/jwt"
)

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = json.GeneratePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
