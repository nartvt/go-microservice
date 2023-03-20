package main

import "auth-component/app/infra/grpc"
import "auth-component/app/infra/db"

func main() {
	db.InitPostgres()
	grpc.InitGrpcServer()
}

func Close() {
	db.ClosePostgres()
}
