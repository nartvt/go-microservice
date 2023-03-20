package main

import (
	"auth-service-client/app/config"
	tClient "auth-service-client/app/transport/grpc/client"
	"auth-service/app/infra/grpc"
)

func main() {
	setupInfra()
	defer closeInfra()
}

func setupInfra() {
	// Every client connect to grpc server
	// must build their own configuration themselves.
	conf := config.Config
	grpcConf := grpc.GrpcConfigClient{
		Host:         conf.AuthGrpcConfig.Host,
		Port:         conf.AuthGrpcConfig.Port,
		ReadTimwOut:  conf.AuthGrpcConfig.ReadTimeout,
		WriteTimeOut: conf.AuthGrpcConfig.WriteTimeout,
	}
	grpc.InitGrpcClient(grpcConf)
	pingGrpc()
}

func pingGrpc() {
	tClient.Hello.Ping()
}

func closeInfra() {
	grpc.CloseGrpcClient()
}
