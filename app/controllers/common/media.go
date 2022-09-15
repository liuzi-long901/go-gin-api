package common

import (
	"github.com/gin-gonic/gin"
	"jassue-gin/app/common/request"
	"jassue-gin/app/common/response"
	"jassue-gin/app/services"
)

// ImageUpload
// Info 文件上传
// @Summary 文件上传
// @Tags 文件上传
// @Accept  json
// @Produce  json
// @Param json body request.ImageUpload true "object json"
// @Success 200 {object} services.OutPut
// @Router /api/auth/upload [post]
func ImageUpload(c *gin.Context) {
	var form request.ImageUpload
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	outPut, err := services.MediaService.SaveImage(form)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, outPut)
}
