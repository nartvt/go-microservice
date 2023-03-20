package main

import (
	"auth-service/app/infra/grpc/client"
	tClient "auth-service/app/transport/grpc/client"
)

func main() {
	setupInfra()
	defer closeInfra()
}

func setupInfra() {
	client.InitGrpcClient()
	pingGrpc()
}

func pingGrpc() {
	tClient.Hello.Ping()
}

func closeInfra() {
	client.CloseGrpcClient()
}
