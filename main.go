package main

import (
	"campus_pharmacy_share/config"
	"campus_pharmacy_share/models"
	"campus_pharmacy_share/routers"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// 初始化配置
	config.InitConfig()

	// 初始化数据库
	models.InitDB()

	// 初始化路由
	router := routers.SetupRouter()

	appPort := viper.GetString("app.port")
	if appPort == "" {
		appPort = "8081"
	}

	fmt.Printf("服务器运行在 http://localhost:%s/\n", appPort)
	err := router.Run(":" + appPort)
	if err != nil {
		panic("服务器运行失败: " + err.Error())
	}
}
