package main

import "auth-service/app/infra/grpc"
import "auth-service/app/infra/db"

func main() {
	db.InitPostgres()
	grpc.InitGrpcServer()
}

func Close() {
	db.ClosePostgres()
}
