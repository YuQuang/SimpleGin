package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/royxu/simplegin/v2/internal/service"
)

type UserController struct {
	UserService *service.UserService
}

type CreateUserRequest struct {
	Username string `json:"username" example:"roy" binding:"required"`
	Email    string `json:"email" example:"a@b.com" binding:"required"`
	Password string `json:"password" example:"password123" binding:"required"`
}

// @Summary 創建用戶
// @Tags User
// @version 1.0
// @Param request body CreateUserRequest true "create user request"
// @produce json
// @Success 200
// @Router /users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request body",
		})
		return
	}

	err := uc.UserService.CreateUser(
		req.Email,
		req.Username,
		req.Password,
	)
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

// @Summary 刪除用戶
// @Tags User
// @version 1.0
// @Param	id	path	int	true	"user search by id"
// @produce json
// @Success 200
// @Router /users/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "failed to get user id",
		})
		return
	}

	err = uc.UserService.DeleteUser(id)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "failed to delete user",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

// @Summary 獲取用戶信息
// @Tags User
// @version 1.0
// @Param	id	path	int	true	"user search by id"
// @produce json
// @Success 200
// @Router /users/{id} [get]
func (uc *UserController) GetUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
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

// @Summary 獲取所有用戶信息
// @Tags User
// @version 1.0
// @produce json
// @Success 200
// @Router /users [get]
func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.UserService.GetUsers()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "failed to get users",
		})
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

// @Summary 更新用戶信息
// @Tags User
// @version 1.0
// @Param	id	path	int	true	"user search by id"
// @Param	username	query	string	false	"username"
// @Param	email	query	string	false	"email"
// @produce json
// @Success 200
// @Router /users/{id} [patch]
func (uc *UserController) PatchUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid id",
		})
		return
	}
	err = uc.UserService.PatchUser(
		c.Query("email"),
		c.Query("username"),
		id,
	)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Patch user failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

// @Summary 覆蓋用戶信息
// @Tags User
// @version 1.0
// @Param	id	path	int	true	"user search by id"
// @Param	username	query	string	true	"username"
// @Param	email	query	string	true	"email"
// @produce json
// @Success 200
// @Router /users/{id} [put]
func (uc *UserController) PutUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid id",
		})
		return
	}
	email := c.Query("email")
	if email == "" {
		c.JSON(400, gin.H{
			"message": "missing email",
		})
		return
	}
	username := c.Query("username")
	if username == "" {
		c.JSON(400, gin.H{
			"message": "missing username",
		})
		return
	}
	err = uc.UserService.PutUser(
		email,
		username,
		id,
	)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Put user failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}
