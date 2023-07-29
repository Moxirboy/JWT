package usecase

import (
	"JWT/internal/models"
	"fmt"

	json "JWT/pkg/jwt"
)

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = json.GeneratePasswordHash(user.Password)
	fmt.Println(user.Password)
	return s.repo.CreateUsers(user)
}
