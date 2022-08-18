package route

import "github.com/gin-gonic/gin"

type route func(engine *gin.Engine)

var routes []route

var middlewares []gin.HandlerFunc

func InitRoute() *gin.Engine {
	g := gin.Default()
	for _,rt := range routes {
		rt(g)
	}
	return g
}
func RegisterRoutes(route ...route) {
	routes = append(routes,route...)
}
func RegisterMiddleware(middleware ...gin.HandlerFunc) {
	middlewares = append(middlewares,middleware...)
}