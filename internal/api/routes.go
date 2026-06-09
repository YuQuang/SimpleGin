package api

import (
	"github.com/gin-gonic/gin"
	"github.com/royxu/simplegin/v2/internal/app"
)

func SetupRoutes(router *gin.RouterGroup, app *app.App) {

	users := router.Group("/users")
	{
		users.POST("", app.UserController.CreateUser)
		users.GET("", app.UserController.GetUsers)

		users.DELETE("/:id", app.UserController.DeleteUser)
		users.PATCH("/:id", app.UserController.PatchUser)
		users.GET("/:id", app.UserController.GetUser)
	}

}
