package app

import (
	"database/sql"

	"github.com/royxu/simplegin/v2/configs"
	"github.com/royxu/simplegin/v2/internal/controller"
	"github.com/royxu/simplegin/v2/internal/repository"
	"github.com/royxu/simplegin/v2/internal/service"
)

type App struct {
	UserController *controller.UserController
}

func InitApp(config *configs.Configuration, db *sql.DB) App {
	userRepository := &repository.UserRepository{
		DB: db,
	}
	userService := &service.UserService{
		UserRepository: userRepository,
	}
	UserController := &controller.UserController{
		UserService: userService,
	}

	return App{
		UserController: UserController,
	}
}
