package conf

import (
	"log"
	"server/globals"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfigByViper() {

	log.Println("配置文件初始化。。。")

	viper.SetConfigType("yaml")
	viper.SetConfigFile("./conf/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err.Error())
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("配置发生变化：", in.Name)
		if err := viper.Unmarshal(&globals.Confok); err != nil {
			log.Println(err.Error())
		}
	})

	if err := viper.Unmarshal(&globals.Confok); err != nil {
		log.Println(err.Error())
	}

}
