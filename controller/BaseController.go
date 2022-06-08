package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"yishixincheng/grace-gin/response"
)

type Controller struct {
}

func (c Controller) parseParameter(ctx *gin.Context, obj any)  {
	err := ctx.ShouldBind(obj)
	if err != nil {
		log.Println(err.Error())
		response.Fail(ctx, "数据验证错误：" + err.Error(), nil)
		return
	}
}
