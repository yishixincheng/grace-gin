package router

import (
	"github.com/gin-gonic/gin"
	"yishixincheng/grace-gin/middleware"
)

type method struct {
	path string
	handler gin.HandlerFunc
	middleware gin.HandlerFunc
}

type route struct {
	group string
	middlewares []middleware.Middleware
	posts []method
	gets []method
	puts []method
	deletes []method
}

type IRestfulFunc interface {
	GET(string, ...gin.HandlerFunc) gin.IRoutes
	POST(string, ...gin.HandlerFunc) gin.IRoutes
	DELETE(string, ...gin.HandlerFunc) gin.IRoutes
	PUT(string, ...gin.HandlerFunc) gin.IRoutes
}

var routeHandlers []func() *route

func GetRoutes(r *gin.Engine) {
	if routeHandlers == nil || len(routeHandlers) == 0 {
		panic("没有注册路由")
	}
	var routes []*route
	for i := range routeHandlers {
		sliceRoute := routeHandlers[i]()
		if sliceRoute == nil {
			continue
		}
		routes = append(routes, sliceRoute)
	}
	if routes == nil || len(routes) == 0{
		panic("没有注册路由")
	}
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	for i := range routes {
		addRoute(r, routes[i])
	}
}

func addRoute(r *gin.Engine, rt *route) {
	if rt.group == "" {
		addRoutePart(r, rt)
	} else {
		addRoutePart(r.Group(rt.group), rt)
	}
}

func addRoutePart(r IRestfulFunc, rt *route) {
	if rt.posts != nil {
		for i := range rt.posts {
			if rt.posts[i].middleware == nil {
				r.POST(rt.posts[i].path, rt.posts[i].handler)
			} else {
				r.POST(rt.posts[i].path, rt.posts[i].middleware, rt.posts[i].handler)
			}
		}
	}
	if rt.gets != nil {
		for i := range rt.gets {
			if rt.gets[i].middleware == nil {
				r.GET(rt.gets[i].path, rt.gets[i].handler)
			} else {
				r.GET(rt.gets[i].path, rt.gets[i].middleware, rt.gets[i].handler)
			}
        }
	}
	if rt.puts != nil {
		for i := range rt.puts {
			if rt.puts[i].middleware == nil {
				r.PUT(rt.puts[i].path, rt.puts[i].handler)
			} else {
				r.PUT(rt.puts[i].path, rt.puts[i].middleware, rt.puts[i].handler)
			}
		}
    }
	if rt.deletes != nil {
		for i := range rt.deletes {
			if rt.deletes[i].middleware == nil {
				r.DELETE(rt.deletes[i].path, rt.deletes[i].handler)
			} else {
				r.DELETE(rt.deletes[i].path, rt.deletes[i].middleware, rt.deletes[i].handler)
			}
		}
	}
}

func registerRouteFunc(fun func() *route) {
	routeHandlers = append(routeHandlers, fun)
}
