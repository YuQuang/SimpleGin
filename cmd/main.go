package main

import (
	"github.com/gin-gonic/gin"
	"github.com/royxu/simplegin/v2/configs"
	"github.com/royxu/simplegin/v2/internal/api"
	App "github.com/royxu/simplegin/v2/internal/app"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger
// @contact.name roy.xu
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:81
// schemes http
func main() {
	// Initialize configuration
	var config = configs.InitConfig("./configs/config.yaml")

	// Initialize Controllers and Services
	var db = App.InitDB(&config)
	defer db.Close()
	var app = App.InitApp(&config, db)

	router := gin.Default()
	App.InitSwagger(&router.RouterGroup, &config)
	api.SetupRoutes(&router.RouterGroup, &app)
	router.Run(config.ServerHost + ":" + config.ServerPort)
}
