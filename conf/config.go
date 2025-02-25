package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	AppName    string
	AppHost    string
	AppPort    string
	AppVersion string
}

func InitConf(configFilePath string) (*AppConfig, error) {
	viper.SetConfigName("conf") // 设置配置文件名（不带扩展名）
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if ok {
			log.Println("找不到配置文件。。。")
			return nil, err
		} else {
			log.Println("配置文件出错。。。")
			return nil, err
		}
	}

	config := &AppConfig{
		AppName:    viper.GetString("app.name"),
		AppHost:    viper.GetString("app.host"),
		AppPort:    viper.GetString("app.port"),
		AppVersion: viper.GetString("app.version"),
	}

	fmt.Println("配置文件加载成功:", config)
	return config, nil
}
