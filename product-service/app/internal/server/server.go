package main

import "product-service/app/infra/grpc"
import "product-service/app/infra/db"

func main() {
	db.InitPostgres()
	grpc.InitGrpcServer()
}

func Close() {
	db.ClosePostgres()
}
