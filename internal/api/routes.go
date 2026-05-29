package api

import (
	"github.com/gin-gonic/gin"

	"github.com/royxu/simplegin/v2/internal/app"
)

func SetupRoutes(router *gin.RouterGroup, app *app.App) {
	router.GET("/json", app.UserController.GetUser)
}
