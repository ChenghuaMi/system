package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"system/global"
)

const configFile = "E:\\go\\src\\system\\config\\config.yml"
func Viper() *viper.Viper {
	vip := viper.New()
	vip.SetConfigFile(configFile)
	vip.SetConfigType("yml")
	if err := vip.ReadInConfig();err != nil {
		panic(fmt.Sprint("fail error %s \n",err.Error()))
	}
	vip.WatchConfig()
	vip.OnConfigChange(func(in fsnotify.Event) {
		loadConfig(vip)
	})
	loadConfig(vip)
	return vip
}
func loadConfig(vip *viper.Viper) {
	vip.Unmarshal(&global.Config)
}
