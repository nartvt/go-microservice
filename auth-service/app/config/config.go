package config

import (
	"github.com/spf13/viper"
	"log"
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
	ConnectTimeOut  int
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
	Host         string
	Port         int
	UserName     string
	Password     string
	Database     string
	MaxIdleConns int
	MaxOpenConns int
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
		log.Panic("Config can't read", err.Error())
		return
	}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Panic("Config can't load", err.Error())
		return
	}
	log.Println(conf)
}
