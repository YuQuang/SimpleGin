package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/royxu/simplegin/v2/internal/service"
	"github.com/royxu/simplegin/v2/internal/utils"
)

type AuthController struct {
	JWTManager  *utils.JWTManager
	AuthService *service.AuthService
}
type LoginRequest struct {
	Username string `json:"username" example:"roy"`
	Email    string `json:"email" example:"a@b.com"`
	Password string `json:"password" example:"password123"`
}

// @Summary 用戶登入
// @Tags Auth
// @version 1.0
// @Param	username	query	string	true	"login username"
// @Param	email	query	string	true	"login email"
// @Param	password	query	string	true	"login password"
// @produce json
// @Success 200
// @Router /login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request body",
		})
		return
	}

	token, err := ac.AuthService.Login(req.Email, req.Username, req.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "login failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "ok",
		"token":   token,
	})
}
