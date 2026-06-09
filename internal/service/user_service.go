package service

import (
	"github.com/royxu/simplegin/v2/internal/model"
	"github.com/royxu/simplegin/v2/internal/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (us *UserService) CreateUser(
	email string,
	username string,
) error {

	err := us.UserRepository.CreateUser(&model.User{
		Username: username,
		Email:    email,
	})
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) DeleteUser(id int) error {
	return us.UserRepository.DeleteUser(id)
}

func (us *UserService) GetUser(
	id int,
) (*model.User, error) {
	user, err := us.UserRepository.GetUser(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) GetUsers() (*[]model.User, error) {
	users, err := us.UserRepository.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserService) PatchUser(
	email string,
	username string,
	id int,
) error {
	err := us.UserRepository.PatchUser(&model.User{
		Email:    email,
		Username: username,
		ID:       int64(id),
	})

	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) PutUser(
	email string,
	username string,
	id int,
) error {
	err := us.UserRepository.PutUser(&model.User{
		Email:    email,
		Username: username,
		ID:       int64(id),
	})

	if err != nil {
		return err
	}

	return nil
}
