package config

import (
	"github.com/spf13/viper"
	"yishixincheng/grace-gin/pkg/dbdriver"
	"os"
)

// InitConfig 初始化配置
func InitConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func InitDefaultDbConnect() {
	driverName := viper.GetString("datasource.driver")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")

	config := dbdriver.ConnectConfig{
		Name: "default",
		Host: host,
		Port: port,
		Driver: driverName,
		Database: database,
		Password: password,
		Username: username,
		Charset: charset,
		Loc: loc,
	}
	dbdriver.OpenConnect(config)
}

func RootDir() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	return wd
}