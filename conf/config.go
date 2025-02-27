package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type AppConfig struct {
	System system `yaml:"system"`
	Logger logger `yaml:"logger"`
}

type system struct {
	AppName    string `yaml:"name"`
	AppHost    string `yaml:"host"`
	AppPort    string `yaml:"port"`
	AppVersion string `yaml:"version"`
	Env        string `yaml:"env"`
}

type logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
}

var Config *AppConfig

func init() {
	yamlFile, err := ioutil.ReadFile("./conf/conf.yaml")
	if err != nil {
		fmt.Println("yamlFile.Get config err ")
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}

}
