package api

import (
	"github.com/gin-gonic/gin"
	"github.com/royxu/simplegin/v2/internal/app"
)

func SetupRoutes(router *gin.RouterGroup, app *app.App) {

	router.GET("/user", app.UserController.GetUser)
	router.POST("/user", app.UserController.CreateUser)
}
