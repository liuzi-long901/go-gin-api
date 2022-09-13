package main

import (
	"github.com/gin-gonic/gin"
	"jassue-gin/bootstrap"
	"jassue-gin/global"
	"net/http"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	r := gin.Default()

	// 初始化日志
	bootstrap.InitZapLogger()

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	// 初始化验证器
	bootstrap.InitializeValidator()
	// 启动服务器
	bootstrap.RunServer()

}
