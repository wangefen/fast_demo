package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		// Host     string
		// Port     string
		// User     string
		// Password string
		// Name     string
		Dsn          string `yaml:"dsn"` //该字段对应yaml中的dsn
		MaxIdleConns int
		MaxOpenCons  int
	}
}

var AppConfig *Config

func InitConfig() {
	//viper在读取yaml文件信息
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	//viper自己读取完yaml信息后，自己将信息匹配到相应的结构体变量中去
	//这样的前提是yaml中的变量字段名和结构体中的字段名必须一样。否则找不到
	//若不想一模一样，可以用yaml标签
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	initDB()
}
