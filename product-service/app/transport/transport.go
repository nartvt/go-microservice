package transport

import (
	"context"
	"fmt"
	"product-service/app/infra/grpc"
	"product-service/app/proto-gen/message"
)

var ProductTransport *productTransport

type productTransport struct{}

func init() {
	ProductTransport = &productTransport{}
}

func (productTransport) GetProductById(productId int) (*message.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), grpc.GetProductGrpcReadTimeout())
	defer cancel()
	req := &message.ProductRequest{
		Id: int64(productId),
	}
	resp, err := grpc.GetGrpcClient().ProductService().GetProductById(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return resp, err
	}
	fmt.Println(resp)
	return resp, nil
}

func (productTransport) GetProducts(productId int, limit int) (*message.ProductResponses, error) {
	ctx, cancel := context.WithTimeout(context.Background(), grpc.GetProductGrpcReadTimeout())
	defer cancel()
	req := &message.ProductRequest{
		Id:    int64(productId),
		Limit: int32(limit),
	}
	resp, err := grpc.GetGrpcClient().ProductService().GetProducts(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return resp, err
	}
	fmt.Println(resp)
	return resp, nil
}
