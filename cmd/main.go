package main

import (
	"github.com/gin-gonic/gin"
	"github.com/royxu/simplegin/v2/configs"
	"github.com/royxu/simplegin/v2/internal/api"
	"github.com/royxu/simplegin/v2/internal/app"
)

func main() {
	// Initialize configuration
	var config = configs.InitConfig()

	// Initialize Controllers and Services
	var db = app.InitDB(&config)
	var app = app.InitApp(&config, db)

	router := gin.Default()
	api.SetupRoutes(&router.RouterGroup, &app)
	router.Run(config.ServerHost + ":" + config.ServerPort)
}
