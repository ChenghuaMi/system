package manager

import (
	"github.com/gin-gonic/gin"
	form2 "system/form"
	"system/form/manager"
	"system/response"
	"system/service"
)

func ListAdmin(ctx *gin.Context) {
	var form form2.Page
	if err := ctx.Bind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	obj,err := service.BaseServiceObj.AdminServiceObj.List(form)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",obj))
	return
}
func AddAdmin(ctx *gin.Context) {
	var form manager.AddAdminForm
	if err := ctx.Bind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	obj,err := service.BaseServiceObj.AdminServiceObj.AddAdmin(form)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",obj))
	return
}
func EditAdmin(ctx *gin.Context) {
	var form manager.EditAdminForm
	if err := ctx.Bind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	obj,err := service.BaseServiceObj.AdminServiceObj.EditAdmin(form)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",obj))
	return
}
func DeleteAdmin(ctx *gin.Context) {
	var form manager.AdminForm
	if err := ctx.Bind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	err := service.BaseServiceObj.AdminServiceObj.DeleteAdmin(form)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",nil))
	return
}