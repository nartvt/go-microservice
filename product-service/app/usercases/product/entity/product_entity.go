package entity

import (
	"context"
	"fmt"
	"product-service/app/domain/model"
	"product-service/app/domain/usercases/product/repository"
	"product-service/app/infra/db"
	"product-service/app/proto-gen/message"
	"product-service/app/proto-gen/rpc"
	"product-service/app/uerror"
	"time"

	"gorm.io/gorm"
)

type productEntity struct {
	rpc.UnimplementedProductServiceServer
}

var ProductEntity rpc.ProductServiceServer

func init() {
	ProductEntity = &productEntity{
		UnimplementedProductServiceServer: rpc.UnimplementedProductServiceServer{},
	}
}

func (p productEntity) GetProductById(ctx context.Context, request *message.ProductRequest) (*message.ProductResponse, error) {
	if request == nil || request.Id <= 0 {
		return &message.ProductResponse{}, nil
	}
	product, err := repository.ProductRepository.GetProductById(int(request.Id))
	if err != nil && err == gorm.ErrRecordNotFound {
		return &message.ProductResponse{}, nil
	}
	if err != nil {
		return &message.ProductResponse{}, uerror.InternalError(err, err.Error())
	}
	if product == nil {
		return &message.ProductResponse{}, uerror.BadRequestError(fmt.Errorf("promotion not found %d", request.Id), "promotion not found")
	}
	return p.bind(product), nil
}

func (p productEntity) GetProducts(ctx context.Context, request *message.ProductRequest) (*message.ProductResponses, error) {
	if request == nil {
		return &message.ProductResponses{}, nil
	}
	if request.Id <= 0 && request.Limit <= 0 {

	}
	products, err := repository.ProductRepository.GetProductsPagination(int(request.Id), int(request.Limit))
	if err != nil {
		return &message.ProductResponses{}, uerror.InternalError(err, err.Error())
	}
	if len(products) <= 0 {
		return &message.ProductResponses{}, nil
	}
	return p.binds(products), nil
}

func (p productEntity) UpdateProduct(crx context.Context, request *message.ProductRequest) (*message.ProductResponse, error) {
	if request == nil || request.Id <= 0 {
		return &message.ProductResponse{}, nil
	}
	product, err := repository.ProductRepository.GetProductById(int(request.Id))
	if err != nil && err == gorm.ErrRecordNotFound {
		return &message.ProductResponse{}, nil
	}
	if err != nil {
		return &message.ProductResponse{}, uerror.InternalError(err, err.Error())
	}
	if product == nil {
		return &message.ProductResponse{}, uerror.BadRequestError(fmt.Errorf("promotion not found %d", request.Id), "promotion not found")
	}
	if request.Active != nil {
		product.Active = request.Active.Active
	}
	if len(request.Image) > 0 {
		product.Image = request.Image
	}
	if request.Price > 0 {
		product.Price = request.Price
	}
	if len(request.Name) > 0 {
		product.Name = request.Name
	}
	tx := db.BeginTx()
	defer db.RecoveryTx(tx)

	err = repository.ProductRepository.UpdateProductTx(product, tx)
	if err != nil {
		tx.Rollback()
		return &message.ProductResponse{}, uerror.InternalError(err, err.Error())
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return &message.ProductResponse{}, uerror.InternalError(err, err.Error())
	}
	return p.bind(product), nil
}

func (p productEntity) CreateProduct(ctx context.Context, newProduct *message.ProductRequest) (*message.ProductResponse, error) {
	if newProduct == nil {
		return &message.ProductResponse{}, nil
	}

	tx := db.BeginTx()
	defer db.RecoveryTx(tx)

	now := time.Now()
	newProductEntity := &model.Product{
		Name:      newProduct.Name,
		Active:    true,
		Price:     newProduct.Price,
		Image:     newProduct.Image,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	err := repository.ProductRepository.CreateProductTx(newProductEntity, tx)
	if err != nil {
		tx.Rollback()
		return &message.ProductResponse{}, uerror.InternalError(err, err.Error())
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return &message.ProductResponse{}, uerror.InternalError(err, err.Error())
	}
	return p.bind(newProductEntity), nil
}

func (productEntity) bind(product *model.Product) *message.ProductResponse {
	if product == nil {
		return &message.ProductResponse{}
	}
	return &message.ProductResponse{
		Id:    int64(product.Id),
		Name:  product.Name,
		Image: product.Image,
		Active: &message.Active{
			Active: product.Active,
		},
		Price:     product.Price,
		CreatedAt: product.CreatedAt.UnixMilli(),
		UpdatedAt: product.UpdatedAt.UnixMilli(),
	}
}

func (p productEntity) binds(products []model.Product) *message.ProductResponses {
	if len(products) <= 0 {
		return &message.ProductResponses{}
	}
	resp := make([]*message.ProductResponse, len(products))
	for i := range products {
		resp[i] = p.bind(&products[i])
	}
	return &message.ProductResponses{
		ProductResponses: resp,
	}
}
