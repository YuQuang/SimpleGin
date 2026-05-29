package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/royxu/simplegin/v2/internal/service"
)

type UserController struct {
	UserService *service.UserService
}

func (uc *UserController) GetUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid id",
		})
		return
	}

	user, err := uc.UserService.GetUser(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "user not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"username": user.Username,
		"email":    user.Email,
	})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	type CreateUserRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request body",
		})
		return
	}

	err := uc.UserService.CreateUser(req.Email, req.Username)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "ok",
	})
}
