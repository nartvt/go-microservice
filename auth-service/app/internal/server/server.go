package main

import (
	"auth-service/app/infra/grpc"
)
import "auth-service/app/infra/db"

func main() {
	setupInfra()
}

func setupInfra() {
	db.InitPostgres()
	defer Close()
	grpc.InitGrpcServer()
}

func Close() {
	db.ClosePostgres()
}
