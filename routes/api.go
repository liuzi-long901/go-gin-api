package routes

import (
	"github.com/gin-gonic/gin"
	"jassue-gin/app/controllers/app"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", app.Register)
}
