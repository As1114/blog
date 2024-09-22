package core

import (
	"fmt"
	"github.com/axis1114/blog/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

const filePath = "settings.yaml"

func InitConf() {
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {
		log.Fatalf("读取配置信息失败, err:%v\n", err)
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		log.Fatalf("viper反序列化失败, err:%v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生改变...")
		if err := viper.Unmarshal(global.Config); err != nil {
			log.Fatalf("viper反序列化失败, err:%v\n", err)
		}
	})
	return
}
