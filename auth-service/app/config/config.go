package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config config

type config struct {
	Posgres        DBConnectInfo
	Authorization  auth
	AuthGrpcConfig authGrpcConfig
}

type auth struct {
	SecretKey  string
	ExpireTime int
}

type authGrpcConfig struct {
	Host         string
	Port         int
	ReadTimeout  int
	WriteTimeout int
}

type DBConnectInfo struct {
	Host     string
	Port     int
	UserName string
	Password string
	Database string
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
