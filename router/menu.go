package router

import "yishixincheng/grace-gin/controller"

func init() {
	registerRouteFunc(func() *route {
		menuController := controller.NewMenuController()
		return &route{
			group: "/menu",
			middlewares: nil,
			posts: []method {
				{
					path: "/list",
					handler: menuController.MenuList,
				},
			},
		}
	})

}
