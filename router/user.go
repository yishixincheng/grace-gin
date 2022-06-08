package router

import "yishixincheng/grace-gin/controller"

func init() {
	registerRouteFunc(func() *route {
		userController := controller.NewUserController()
		return &route{
			group: "/user",
			middlewares: nil,
			posts: []method {
				{
					path: "/login",
					handler: userController.Login,
				},
				{
					path: "/create",
					handler: userController.Create,
				},
			},
			gets: []method {
				{
					path: "hello",
					handler: userController.Hello,
				},
			},
		}
	})
}