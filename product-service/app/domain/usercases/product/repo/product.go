package repo

import (
	"context"
	"fmt"
	"product-service/app/domain/entities"
	"product-service/app/domain/usercases/product/orm"
	"product-service/app/infra/db"
	"product-service/app/transport/proto-gen/message"
	"product-service/app/transport/proto-gen/rpc"
	"product-service/app/uerror"
	"time"

	"gorm.io/gorm"
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
	if request == nil || request.Id <= 0 {
		return &message.ProductResponse{}, nil
	}
	product, err := orm.Product.GetProductById(int(request.Id))
	if err != nil && err == gorm.ErrRecordNotFound {
		return &message.ProductResponse{}, nil
	}
	if err != nil {
		return &message.ProductResponse{}, uerror.InternalError(err, err.Error())
	}
	if product == nil {
		return &message.ProductResponse{}, uerror.BadRequestError(fmt.Errorf("product not found %d", request.Id), "product not found")
	}
	return p.bind(product), nil
}
func (p productRepo) GetProducts(ctx context.Context, request *message.ProductRequest) (*message.ProductResponses, error) {
	if request == nil {
		return &message.ProductResponses{}, nil
	}
	if request.Id <= 0 && request.Limit <= 0 {

	}
	products, err := orm.Product.GetProductsPagination(int(request.Id), int(request.Limit))
	if err != nil {
		return &message.ProductResponses{}, uerror.InternalError(err, err.Error())
	}
	if len(products) <= 0 {
		return &message.ProductResponses{}, nil
	}
	return p.binds(products), nil
}

func (p productRepo) UpdateProduct(crx context.Context, request *message.ProductRequest) (*message.ProductResponse, error) {
	if request == nil || request.Id <= 0 {
		return &message.ProductResponse{}, nil
	}
	product, err := orm.Product.GetProductById(int(request.Id))
	if err != nil && err == gorm.ErrRecordNotFound {
		return &message.ProductResponse{}, nil
	}
	if err != nil {
		return &message.ProductResponse{}, uerror.InternalError(err, err.Error())
	}
	if product == nil {
		return &message.ProductResponse{}, uerror.BadRequestError(fmt.Errorf("product not found %d", request.Id), "product not found")
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

	err = orm.Product.UpdateProductTx(product, tx)
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

func (p productRepo) CreateProduct(ctx context.Context, newProduct *message.ProductRequest) (*message.ProductResponse, error) {
	if newProduct == nil {
		return &message.ProductResponse{}, nil
	}

	tx := db.BeginTx()
	defer db.RecoveryTx(tx)

	now := time.Now()
	newProductEntity := &entities.Product{
		Name:      newProduct.Name,
		Active:    true,
		Price:     newProduct.Price,
		Image:     newProduct.Image,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	err := orm.Product.CreateProductTx(newProductEntity, tx)
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

func (productRepo) bind(product *entities.Product) *message.ProductResponse {
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

func (p productRepo) binds(products []entities.Product) *message.ProductResponses {
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
