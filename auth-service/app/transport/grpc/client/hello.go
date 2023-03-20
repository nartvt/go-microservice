package client

import (
	"auth-component/app/proto-gen/message"
	"auth-service/app/infra/grpc/client"
	"context"
	"fmt"
)

var Hello *hello

type hello struct{}

func init() {
	Hello = &hello{}
}

func (hello) Ping() (*message.HelloResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), client.GetAuthGrpcReadTimeout())
	defer cancel()
	req := &message.HelloRequest{Name: "From Client: Ping"}
	resp, err := client.GrpcClient().Hello().SayHello(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return resp, err
	}
	fmt.Println(resp)
	return resp, nil
}
