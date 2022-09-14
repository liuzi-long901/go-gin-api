package app

import (
	"github.com/gin-gonic/gin"
	"jassue-gin/app/common/request"
	"jassue-gin/app/common/response"
	"jassue-gin/app/services"
)

// Register 用户注册
// @Summary 用户注册
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param json body request.Register true "object json"
// @Success 200 {object} models.User
// @Router /api/auth/register [post]
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
