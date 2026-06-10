package service

import (
	"fmt"

	"github.com/royxu/simplegin/v2/internal/repository"
	"github.com/royxu/simplegin/v2/internal/utils"
)

type AuthService struct {
	JWTManager     *utils.JWTManager
	UserRepository *repository.UserRepository
}

func (as *AuthService) Login(
	identifier string,
	password string,
) (string, error) {

	user, err := as.UserRepository.GetUserByIdentifier(identifier)
	if err != nil {
		return "", err
	}

	res := utils.VerifyPassword(user.Password, password)
	if !res {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := as.JWTManager.GenerateToken(
		uint(user.ID),
		user.Username,
		user.Email,
	)
	if err != nil {
		return "", err
	}

	return token, nil
}
