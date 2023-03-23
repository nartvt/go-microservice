package repo

import (
	"context"
	"product-service/app/domain/entities"
	"product-service/app/domain/models"
	"product-service/app/domain/usercases/common"
	"product-service/app/domain/usercases/product/orm"
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

func (productRepo) GetProductById(ctx context.Context, request *message.ProductRequest) (*message.ProductResponse, error) {
	_, err := orm.Product.GetProductById(int(request.Id))
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, uerror.InternalError(err, err.Error())
	}
	return nil, nil
}

func (productRepo) UpdateProduct(crx context.Context, product *message.ProductRequest) (*message.ProductResponse, error) {
	return nil, nil
}

func (productRepo) CreateProductTx(newProduct models.ProductRepo) (models.ProductRepo, error) {
	tx := common.BeginTx()
	defer common.RecoveryTx(tx)

	now := time.Now()
	newProductEntity := entities.Product{
		Name:      newProduct.Name,
		Active:    true,
		Image:     newProduct.Image,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	err := orm.Product.CreateProductTx(&newProductEntity, tx)
	if err != nil {
		tx.Rollback()
		return models.ProductRepo{}, uerror.InternalError(err, err.Error())
	}
	return models.ProductRepo{
		Id:        newProductEntity.Id,
		Active:    newProductEntity.Active,
		Image:     newProductEntity.Image,
		CreatedAt: newProductEntity.CreatedAt,
		UpdatedAt: newProductEntity.UpdatedAt,
	}, nil
}

func (d productRepo) populateProduct(product *entities.Product) *models.ProductRepo {
	if product == nil {
		return nil
	}
	return &models.ProductRepo{
		Id:        product.Id,
		Name:      product.Name,
		Image:     product.Image,
		Active:    product.Active,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}
