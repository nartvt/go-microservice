package main

import (
	"product-service/app/infra/db"
	"product-service/app/infra/grpc"
)

func main() {
	setupInfra()
	defer CloseInfra()
}

func CloseInfra() {
	db.ClosePostgres()
}

func setupInfra() {
	db.InitPostgres()
	grpc.InitGrpcServer()
}
