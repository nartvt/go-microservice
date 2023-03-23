package client

import (
	"auth-service/app/infra/grpc"
	"auth-service/app/proto-gen/message"
	"context"
	"log"
)

var User *user

type user struct{}

func init() {
	User = &user{}
}

func (user) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), grpc.GetAuthGrpcReadTimeout())
	defer cancel()
	req := &message.UserRequest{Name: "From Client: Ping"}
	resp, err := grpc.GetGrpcClient().UserClient().Ping(ctx, req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(resp)
}
