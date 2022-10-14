package main

import (
	"github.com/gin-gonic/gin"
	"github.com/x-funs/go-fun"
	"jassue-gin/bootstrap"
	"jassue-gin/global"
	"jassue-gin/setup"
	"net/http"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	r := gin.Default()

	// 初始化日志
	bootstrap.InitZapLogger()

	////测试send
	//bootstrap.Product()
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
	// 初始化Redis
	global.App.Redis = bootstrap.InitializeRedis()
	//初始化mq

	//初始化文件上传服务 支持本地 阿里云 七牛云
	bootstrap.InitializeStorage()

	//初始化mq
	setup.Rabbit()
	if simple, err := bootstrap.NewSimple("comment"); err == nil {
		for i := 0; i < 100; i++ {
			msg := Demo{
				Id:   i + 1,
				Name: fun.RandomLetter(4),
				Time: fun.Date(fun.DatetimeMilliPattern),
			}
			msgJson := fun.ToJson(msg)
			if err := simple.Send(fun.Bytes(msgJson)); err == nil {
				bootstrap.Info("Send simple success", bootstrap.String("msg", fun.ToString(msg)))
			} else {
				bootstrap.Error("Send simple error")
			}
		}
	}

	// 启动服务器
	bootstrap.RunServer()

}

type Demo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
	Time string `json:"time"`
}
