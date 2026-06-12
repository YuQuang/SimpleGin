package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/royxu/simplegin/v2/internal/service"
)

type MemoController struct {
	MemoService *service.MemoService
}
type CreateMemoRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	UserID   int    `json:"user_id" binding:"required"`
	IsPublic bool   `json:"is_public"`
}

// @Summary 創建一篇 Memo
// @Tags Memo
// @version 1.0
// @Param	title		query	string	true	"Memo title"
// @Param	content		query	string	true	"Memo content"
// @Param	user_id		query	int		true	"Who is owner"
// @Param	is_public	query	bool	true	"Is public or not"
// @produce json
// @Success 200
// @Router /memos [post]
func (mc *MemoController) CreateMemo(c *gin.Context) {
	var req CreateMemoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	memo, err := mc.MemoService.CreateMemo(
		req.Title,
		req.Content,
		req.UserID,
		req.IsPublic,
	)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": memo,
	})
}

// @Summary 取得特定的 Memo
// @Tags Memo
// @version 1.0
// @Param	id		path	string	true	"Memo id"
// @produce json
// @Success 200
// @Router /memos/{id} [get]
func (mc *MemoController) GetMemo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid memo id",
		})
		return
	}

	memo, err := mc.MemoService.GetMemo(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": memo,
	})
}

// @Summary 根據給的 limit 跟 offset 回傳 Memos
// @version 1.0
// @Param	limit		query	string	true	"search limit"
// @Param	offset		query	string	true	"search offset"
// @produce json
// @Success 200
// @Router /memos/{id} [get]
func (mc *MemoController) GetMemos(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid limit",
		})
		return
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid offset",
		})
		return
	}

	memos, err := mc.MemoService.GetMemos(limit, offset)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": memos,
	})
}
