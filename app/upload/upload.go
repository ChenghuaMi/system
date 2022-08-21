package upload

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
	"system/response"
	"system/util"
)

func Upload(ctx *gin.Context) {
	file,_ := ctx.FormFile("thumb")
	ps,_ := filepath.Abs("") // 获取项目的根路径
	ss := util.CreateFold(ps)
	dist := util.ReNameFile(filepath.Join(ss,file.Filename))
	err := ctx.SaveUploadedFile(file,dist)
	if err != nil {
		ctx.JSON(500,response.FailResponse(10001,err.Error()))
		return
	}
	ph := util.ReplacePath(ps,dist)
	ctx.JSON(200,response.SuccessResponse(0,"ok",ph))
	return
}