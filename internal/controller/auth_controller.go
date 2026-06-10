package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/royxu/simplegin/v2/internal/service"
)

type AuthController struct {
	AuthService *service.AuthService
}
type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

// @Summary 用戶登入
// @Tags Auth
// @version 1.0
// @Param	identifier	query	string	true	"login identifier"
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

	token, err := ac.AuthService.Login(req.Identifier, req.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "ok",
		"token":   token,
	})
}
