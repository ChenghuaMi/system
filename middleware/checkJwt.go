package middleware

import (
	"github.com/gin-gonic/gin"
	"system/response"
	"fmt"
)

func CheckJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		fmt.Println("token:",token)
		if token == "" {
			fmt.Println("Xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
			ctx.JSON(500,response.FailResponse(10000,"token 不能为空"))
			fmt.Println("????????????????????????????")
			ctx.Abort()
			return
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		}
		clain,err := ParseToken(token)
		if err != nil {
			ctx.JSON(500,response.FailResponse(10000,err.Error()))

			ctx.Abort()
			return
		}
		ctx.Set("admin_id",clain.Id)
		ctx.Next()
	}
}
