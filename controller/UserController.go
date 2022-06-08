package controller

import (
	"github.com/gin-gonic/gin"
	"yishixincheng/grace-gin/response"
)

type UserController struct {

}

func NewUserController() UserController {
	return UserController{}
}

func (c UserController) Login(ctx *gin.Context) {

}

func (c UserController) Create(ctx *gin.Context) {

}

func (c UserController) Hello(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"token": "dafdafda",
	}, "")
}

