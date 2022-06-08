package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"yishixincheng/grace-gin/config"
	"yishixincheng/grace-gin/pkg/dbdriver"
	"yishixincheng/grace-gin/router"
)

func main() {
	config.InitConfig()
	config.InitDefaultDbConnect()
	defer dbdriver.CloseConnects()

	r := gin.Default()
	router.GetRoutes(r)

	port := viper.GetString("server.port")

	if port != "" {
		panic(r.Run(":" + port))
	} else {
		panic(r.Run())
	}
}

