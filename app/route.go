package app

import (
	"github.com/gin-gonic/gin"
	"system/app/login"
	"system/app/manager"
	"system/app/product"
	"system/app/upload"
	"system/middleware"
	"system/route"
)

func InitRoute() {
	route.RegisterRoutes(Login,Manager,Product,Upload)
}
func InitMiddleware() {
	route.RegisterMiddleware(middleware.CheckJwt())
}

func Login(g *gin.Engine) {
	group := g.Group("login")
	{
		group.POST("login",login.Login)
	}
}
func Manager(g *gin.Engine) {
	group := g.Group("manager")
	{
		group.POST("add_role",manager.AddRole)
		group.POST("edit_role",manager.EditRole)
		group.POST("delete_role",manager.DeleteRole)
		group.GET("list_role",manager.ListRole)

		group.GET("list_admin",manager.ListAdmin)
		group.POST("add_admin",manager.AddAdmin)
		group.POST("edit_admin",manager.EditAdmin)
		group.POST("delete_admin",middleware.CheckJwt(),manager.DeleteAdmin)
	}
}
func Product(g *gin.Engine) {
	group := g.Group("product")
	{
		group.GET("list_cate",product.ListCate)
		group.POST("add_cate",product.AddCate)
		group.POST("edit_cate",product.EditCate)
		group.POST("delete_cate",product.DeleteCate)

		group.GET("list_brand",product.ListBrand)
		group.POST("add_brand",product.AddBrand)
		group.POST("edit_brand",product.EditBrand)
		group.POST("delete_brand",product.DeleteBrand)
	}
}
func Upload(g *gin.Engine) {
	group := g.Group("upload")
	{
		group.POST("upload",upload.Upload)
	}
}