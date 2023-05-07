package main

import (
	"elasticsearch-service/app/infra/grpc"
)

func main() {
	setupInfra()
}

func setupInfra() {
	grpc.InitGrpcServer()
}
