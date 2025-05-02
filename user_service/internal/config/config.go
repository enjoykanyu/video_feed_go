package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// InitConfig 初始化配置文件
func InitConfig() {
	viper.AddConfigPath("./internal/config")
	viper.AddConfigPath(".") //多路径查找
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Failed to read config file: %v\nSearch paths: %v\n",
			err, viper.ConfigFileUsed())
		panic(err)
	}

	//监控并重新读取配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
}
