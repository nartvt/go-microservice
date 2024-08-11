package repo

import (
	"context"
	"elasticsearch-service/app/domain/models"
	"elasticsearch-service/app/domain/storage"
	"elasticsearch-service/app/transport/proto-gen/message"
	"elasticsearch-service/app/transport/proto-gen/rpc"
	"elasticsearch-service/app/uerror"
	"elasticsearch-service/app/util"
	"time"
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
	now := time.Now()
	product := models.ProductModel{
		Id:        int(newProduct.Id),
		Name:      newProduct.Name,
		CreatedAt: util.FormatDateTime(now),
		UpdatedAt: util.FormatDateTime(now),
	}
	_, err := storage.ProductElastic.IndexProduct(product)
	if err != nil {
		return &message.ProductResponse{}, uerror.InternalError(err, err.Error())
	}
	return &message.ProductResponse{
		Id:        int64(product.Id),
		Name:      product.Name,
		CreatedAt: now.UnixMilli(),
		UpdatedAt: now.UnixMilli(),
	}, nil
}
