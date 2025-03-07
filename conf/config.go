package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"time"
)

type AppConfig struct {
	System system `yaml:"system"`
	Logger logger `yaml:"logger"`
	Pgsql  pgsql  `yaml:"pgsql"`
	Jwt    jwt    `yaml:"jwt"`
}

type system struct {
	AppName      string        `yaml:"name"`
	AppHost      string        `yaml:"host"`
	AppPort      string        `yaml:"port"`
	AppVersion   string        `yaml:"version"`
	Env          string        `yaml:"env"`
	ReadTimeOut  time.Duration `yaml:"readTimeOut"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	JwtSecret    string        `yaml:"jwtSecret"`
}

type pgsql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
}

type jwt struct {
	TimeOut    time.Duration `yaml:"timeOut"`
	MaxRefresh time.Duration `yaml:"maxRefresh"`
	Realm      string        `yaml:"realm"`
	Key        string        `yaml:"key"`
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

func (s *system) GetJwtSecret() []byte {
	return []byte(s.JwtSecret) // 转换为 []byte
}
