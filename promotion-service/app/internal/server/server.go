package main

import (
	"promotion-service/app/infra/db"
	"promotion-service/app/infra/grpc"
)

func main() {
	setupInfra()
	defer closeInfra()
}

func setupInfra() {
	grpc.InitGrpcServer()
	db.InitPostgres()
}

func closeInfra() {
	db.ClosePostgres()
}
