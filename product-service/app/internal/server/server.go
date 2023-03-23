package main

import "product-service/app/infra/grpc"
import "product-service/app/infra/db"

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
