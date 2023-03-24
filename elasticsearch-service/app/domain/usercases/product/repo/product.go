package repo

import (
	"context"
	"elasticsearch-service/app/transport/proto-gen/message"
	"elasticsearch-service/app/transport/proto-gen/rpc"
)

type productRepo struct {
	rpc.UnimplementedProductServiceServer
}

var Product rpc.ProductServiceServer

func init() {
	Product = &productRepo{
		UnimplementedProductServiceServer: rpc.UnimplementedProductServiceServer{},
	}
}

func (p productRepo) GetProductById(ctx context.Context, request *message.ProductRequest) (*message.ProductResponse, error) {
	return nil, nil
}
func (p productRepo) GetProducts(ctx context.Context, request *message.ProductRequest) (*message.ProductResponses, error) {
	return nil, nil
}

func (p productRepo) UpdateProduct(crx context.Context, request *message.ProductRequest) (*message.ProductResponse, error) {
	return nil, nil
}

func (p productRepo) CreateProduct(ctx context.Context, newProduct *message.ProductRequest) (*message.ProductResponse, error) {
	return nil, nil
}
