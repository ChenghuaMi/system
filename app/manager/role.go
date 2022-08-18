package manager

import (
	"github.com/gin-gonic/gin"
	form2 "system/form"
	"system/form/manager"
	"system/response"
	"system/service"
)

func AddRole(ctx *gin.Context) {
	var form manager.AddRoleForm
	if err := ctx.ShouldBind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	obj,err := service.BaseServiceObj.RoleServiceObj.AddRole(form)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",obj))
	return
}
func EditRole(ctx *gin.Context) {
	//var form manager.EditRoleForm
	var form manager.EditRoleForm
	if err := ctx.ShouldBind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	obj,err := service.BaseServiceObj.RoleServiceObj.EditRole(form)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",obj))
	return

}
func DeleteRole(ctx *gin.Context) {
	var form manager.RoleForm
	if err := ctx.ShouldBind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	err := service.BaseServiceObj.RoleServiceObj.DeleteRole(form)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",""))
	return

}
func ListRole(ctx *gin.Context) {
	var form form2.Page
	if err := ctx.ShouldBind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	obj,err := service.BaseServiceObj.RoleServiceObj.List(form)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",obj))
	return
}
