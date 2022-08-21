package product

import (
	"github.com/gin-gonic/gin"
	form2 "system/form"
	"system/form/product"
	"system/response"
	"system/service"
)

func AddCate(ctx *gin.Context) {
	form := product.CateForm{}
	if err := ctx.Bind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	obj,err := service.BaseServiceObj.CateServiceObj.AddCate(form)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",obj))
	return
}
func EditCate(ctx *gin.Context) {
	form := product.EditCateForm{}
	if err := ctx.Bind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	obj,err := service.BaseServiceObj.CateServiceObj.EditCate(form)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",obj))
	return
}
func DeleteCate(ctx *gin.Context) {
	form := product.DeleteCateForm{}
	if err := ctx.Bind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ok := service.BaseServiceObj.CateServiceObj.DeleteCate(form)
	if !ok {
		ctx.JSON(500,response.FailResponse(10001,"晒出失败"))
		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",nil))
	return
}
func ListCate(ctx *gin.Context) {
	form := form2.SearchForm{}
	if err := ctx.Bind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	obj,err := service.BaseServiceObj.CateServiceObj.List(form)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",obj))
	return
}