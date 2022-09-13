package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"jassue-gin/global"
)

// @title king-gin-api
// @version 1.0
// @schemes http https
// @basePath /
//
// @securityDefinitions.basic  BasicAuth
//
// @securityDefinitions.apikey ApiJwtAuth
// @in header
// @name Jwt-Authorization
// @description Jwt Header
func initSwagger(r *gin.Engine) {
	swaggerGroup := r.Group("/swagger")
	if global.App.Config.Swagger.Auth {
		swaggerGroup.Use(gin.BasicAuth(gin.Accounts{
			global.App.Config.Swagger.Username: global.App.Config.Swagger.Password,
		}))
	}
	{
		swaggerGroup.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

}
