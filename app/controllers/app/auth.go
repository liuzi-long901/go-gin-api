package app

import (
	"github.com/gin-gonic/gin"
	"jassue-gin/app/common/request"
	"jassue-gin/app/common/response"
	"jassue-gin/app/services"
)

// Login 用户登录
// @Summary 用户登录
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param json body request.Login true "object json"
// @Success 200 {object} services.TokenOutPut
// @Router /api/auth/login [post]
func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

// Info 用户信息
// @Summary 用户信息
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.User
// @Router /api/auth/info [post]
func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}
