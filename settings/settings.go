package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("config.yaml") // 指定配置文件
	viper.AddConfigPath(".")           // 指定查找配置文件的路径
	err = viper.ReadInConfig()         // 读取配置信息
	if err != nil {                    // 读取配置信息失败
		fmt.Println("读取配置信息失败:", err)
		return err
	}

	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置问价修改了")
	})
	return err
}