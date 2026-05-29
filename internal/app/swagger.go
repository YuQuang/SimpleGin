package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/royxu/simplegin/v2/configs"

	_ "github.com/royxu/simplegin/v2/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSwagger(router *gin.RouterGroup, config *configs.Configuration) {
	if mode := gin.Mode(); mode == gin.DebugMode {
		url := ginSwagger.URL(
			fmt.Sprintf("http://%s:%s/swagger/doc.json", config.ServerHost, config.ServerPort),
		)
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}
