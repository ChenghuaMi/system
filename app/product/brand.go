package product

import (
	"github.com/gin-gonic/gin"
	form2 "system/form"
	"system/form/product"
	"system/response"
	"system/service"
)
func ListBrand(ctx *gin.Context) {
	var form form2.SearchForm
	if err := ctx.Bind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	obj,err := service.BaseServiceObj.BrandServiceObj.List(form)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",obj))
	return
}
func AddBrand(ctx *gin.Context) {
	var form product.AddBrandForm
	if err := ctx.Bind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	obj,err := service.BaseServiceObj.BrandServiceObj.AddBrand(form)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",obj))
	return
}
func EditBrand(ctx *gin.Context) {
	var form product.EditBrandForm
	if err := ctx.Bind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	obj,err := service.BaseServiceObj.BrandServiceObj.EditBrand(form)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ctx.JSON(200,response.SuccessResponse(0,"ok",obj))
	return
}
func DeleteBrand(ctx *gin.Context) {
	var form product.DeleteBrandForm
	if err := ctx.Bind(&form);err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	service.BaseServiceObj.BrandServiceObj.DeleteBrand(form.BrandIds)
	ctx.JSON(200,response.SuccessResponse(0,"ok",nil))
	return
}
