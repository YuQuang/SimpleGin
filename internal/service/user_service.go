package service

import (
	"github.com/royxu/simplegin/v2/internal/repository"
)

type UserService struct{}

func (us *UserService) GetUser() string {
	return repository.GetUser()
}
