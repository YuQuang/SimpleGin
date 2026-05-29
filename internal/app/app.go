package app

import (
	"database/sql"

	"github.com/royxu/simplegin/v2/configs"
	"github.com/royxu/simplegin/v2/internal/controller"
)

type App struct {
	UserController *controller.UserController
}

func InitApp(config *configs.Configuration, db *sql.DB) App {
	return App{
		UserController: &controller.UserController{},
	}
}
