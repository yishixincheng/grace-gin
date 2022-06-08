package controller

import (
	"github.com/gin-gonic/gin"
	"yishixincheng/grace-gin/response"
	"yishixincheng/grace-gin/service/menuservice"
	"yishixincheng/grace-gin/vo/input/menuvo"
)

type MenuController struct {
	Controller
	service *menuservice.MenuService
}

func NewMenuController() MenuController {
	return MenuController{
		service: menuservice.NewMenuService(),
	}
}

func (c MenuController) MenuList(ctx *gin.Context)  {
	var menuListVo menuvo.IMenuListRequest
	c.parseParameter(ctx, &menuListVo)

	list := c.service.GetMenuList(menuListVo)

	response.Success(ctx, gin.H{ "list": list}, "")

}
