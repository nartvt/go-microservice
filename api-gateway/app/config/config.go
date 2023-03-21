package config

import (
	"log"

	"github.com/spf13/viper"
)

var conf config

type config struct {
	Postgres       DBConnectInfo
	Server         server
	Auth           auth
	AuthGrpcConfig authGrpcConfig
}

type server struct {
	Host            string
	Port            int
	ReadTimeOut     int
	WriteTimeOut    int
	ShutdownTimeOut int
}

type auth struct {
	SecretKey  string
	ExpireTime int
}

type DBConnectInfo struct {
	Host         string
	Port         int
	UserName     string
	Password     string
	Database     string
	MaxIdleConns int
	MaxOpenConns int
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

func Get() config {
	return conf
}

func load() {
	viper.SetConfigName("config")
	viper.AddConfigPath("app/config")
	if err := viper.ReadInConfig(); err != nil {
		log.Panic("Config can't read")
		return
	}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Panic("Config can't load")
		return
	}
	log.Println(conf)
}
