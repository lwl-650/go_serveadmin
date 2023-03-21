package util

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadeYaml(val string) string {

	config := viper.New()
	config.AddConfigPath("./config/")
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}

	res := config.GetString(val) // 读取配置

	// fmt.Println("viper load ini: +++++++++++++++++++++++++++++++++++", res)
	return res
}
