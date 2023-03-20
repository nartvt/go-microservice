package repo

import (
	"context"
	"product-service/app/proto-gen/message"
	"product-service/app/proto-gen/rpc"
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
