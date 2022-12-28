package app

import (
	"jassue-gin/app/common/response"
	"jassue-gin/app/services"

	"github.com/gin-gonic/gin"
)

// Jijin 基金消息
// @Summary 消息通知
// @Tags 消息管理
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Router /api/message/jijin [get]
func Jijin(c *gin.Context) {
	services.MessageService.Jijin()
	response.Success(c, "基金消息推送成功")
}
