package app

import (
	"github.com/gin-gonic/gin"
	"system/app/login"
	"system/app/manager"
	"system/middleware"
	"system/route"
)

func InitRoute() {
	route.RegisterRoutes(Login,Manager)
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