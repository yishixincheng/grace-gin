package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"yishixincheng/grace-gin/response"
)

func RecoveryMiddleware() Middleware {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(c, fmt.Sprint(err), nil)
			}
		}()
		c.Next()
	}
}