package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/royxu/simplegin/v2/internal/service"
)

type UserController struct {
	UserService *service.UserService
}

func (uc *UserController) GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": uc.UserService.GetUser(),
	})
}
