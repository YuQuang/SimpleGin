package app

import (
	"database/sql"

	"github.com/royxu/simplegin/v2/configs"
	"github.com/royxu/simplegin/v2/internal/controller"
	"github.com/royxu/simplegin/v2/internal/repository"
	"github.com/royxu/simplegin/v2/internal/service"
	"github.com/royxu/simplegin/v2/internal/utils"
)

type App struct {
	UserController *controller.UserController
	AuthController *controller.AuthController
}

func InitApp(config *configs.Configuration, db *sql.DB) App {
	userController := &controller.UserController{
		UserService: &service.UserService{
			UserRepository: &repository.UserRepository{
				DB: db,
			},
		},
	}

	authController := &controller.AuthController{
		AuthService: &service.AuthService{
			JWTManager: utils.NewJWTManager(
				config.JWTSecret,
				config.JWTExpiry,
			),
			UserRepository: &repository.UserRepository{
				DB: db,
			},
		},
	}

	return App{
		UserController: userController,
		AuthController: authController,
	}
}
