package routes

import (
	"jassue-gin/app/controllers/app"
	"jassue-gin/app/controllers/app/middleware"
	"jassue-gin/app/controllers/common"
	"jassue-gin/app/services"

	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", app.Register)
	router.POST("/auth/login", app.Login)
	router.GET("/message/jijin", app.Jijin)
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
		authRouter.POST("/auth/upload", common.ImageUpload)
	}

}
