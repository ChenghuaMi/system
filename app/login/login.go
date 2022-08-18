package login

import (
	"github.com/gin-gonic/gin"
	"system/form/login"
	"system/middleware"
	"system/response"
	"system/service"

	"fmt"
)

func Login(ctx *gin.Context) {
	var form login.AdminForm
	if err := ctx.ShouldBind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,"验证错误"))
		return
	}
	obj,err := service.BaseServiceObj.LoginServiceObj.Login(form)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(500,response.FailResponse(10001,err.Error()))

		return
	}
	token,err := middleware.CreateToken(obj.Id)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))

		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",token))

	return
}