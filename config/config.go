package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func InitConfig() {
	// 设置配置文件名
	viper.SetConfigName("config")
	viper.AddConfigPath("./") // 配置文件所在的目录
	// 设置配置文件类型
	viper.SetConfigType("yaml")

	// 设置配置文件搜索路径
	exePath, _ := os.Executable()
	dir := filepath.Dir(exePath)
	println(dir)
	viper.AddConfigPath(dir)

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置文件失败: " + err.Error())
	}
}
