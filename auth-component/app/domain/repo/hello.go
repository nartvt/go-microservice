package repo

import (
	"auth-component/app/proto-gen/message"
	"auth-component/app/proto-gen/rpc"
	"context"
)

type hello struct {
	rpc.UnimplementedExampleServiceServer
}

func NewHello() *hello {
	return &hello{UnimplementedExampleServiceServer: rpc.UnimplementedExampleServiceServer{}}
}

func (*hello) SayHello(ctx context.Context, req *message.HelloRequest) (*message.HelloResponse, error) {
	return &message.HelloResponse{
		Message: "From Server: Pong",
	}, nil
}
