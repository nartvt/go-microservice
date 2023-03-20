package client

import (
	"auth-service/app/infra/grpc"
	"auth-service/app/proto-gen/message"
	"context"
	"fmt"
)

var Hello *hello

type hello struct{}

func init() {
	Hello = &hello{}
}

func (hello) Ping() (*message.HelloResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), grpc.GetAuthGrpcReadTimeout())
	defer cancel()
	req := &message.HelloRequest{Name: "From Client: Ping"}
	resp, err := grpc.GrpcClient().Hello().SayHello(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return resp, err
	}
	fmt.Println(resp)
	return resp, nil
}
