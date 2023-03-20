package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config config

type config struct {
	AuthGrpcConfig authGrpcConfig
}

type authGrpcConfig struct {
	Host         string
	Port         int
	ReadTimeout  int
	WriteTimeout int
}

func init() {
	load()
}

func load() {
	viper.SetConfigName("config")
	viper.AddConfigPath("app/config")
	if err := viper.ReadInConfig(); err != nil {
		log.Panic("Config can't read")
		return
	}
	if err := viper.Unmarshal(&Config); err != nil {
		log.Panic("Config can't load")
		return
	}
	log.Println(Config)
}
