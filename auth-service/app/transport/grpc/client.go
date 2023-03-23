package grpc

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

func (user) Ping() (*message.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), grpc.GetAuthGrpcReadTimeout())
	defer cancel()
	req := &message.UserRequest{Name: "From Client: Ping"}
	resp, err := grpc.GetGrpcClient().UserClient().Ping(ctx, req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return resp, nil
}

func (user) GetUserByUserName(userName string) (*message.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), grpc.GetAuthGrpcReadTimeout())
	defer cancel()
	req := &message.UserRequest{UserName: userName}
	resp, err := grpc.GetGrpcClient().UserClient().GetUserByUserName(ctx, req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return resp, nil
}
