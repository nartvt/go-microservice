package repository

import (
	"gorm.io/gorm"
	"product-service/app/domain/model"
	"product-service/app/infra/db"
)

type productRepository struct{}

type IProductRepository interface {
	CreateProductTx(newProduct *model.Product, tx *gorm.DB) error
	GetProductById(productId int) (*model.Product, error)
	UpdateProductTx(newProduct *model.Product, tx *gorm.DB) error
	GetProductsPagination(nextId int, limit int) ([]model.Product, error)
}

func (d productRepository) GetProductsPagination(nextId int, limit int) ([]model.Product, error) {
	var resp []model.Product
	query := db.DB().Model(&model.Product{}).
		Where("promotion.active = TRUE").
		Where("deleted_at IS NULL").
		Where("active = ?", true).
		Order("id DESC").
		Limit(limit)
	if nextId > 0 {
		query = query.Where("id < ?", nextId)
	}
	err := query.Find(&resp).Error
	return resp, err
}

var ProductRepository IProductRepository

func init() {
	ProductRepository = productRepository{}
}
func (d productRepository) CreateProductTx(newProduct *model.Product, tx *gorm.DB) error {
	return tx.Create(newProduct).Error
}

func (d productRepository) UpdateProductTx(newProduct *model.Product, tx *gorm.DB) error {
	return tx.Save(newProduct).Error
}

func (d productRepository) GetProductById(productId int) (*model.Product, error) {
	var resp *model.Product
	err := db.DB().Model(&model.Product{}).
		Where("promotion.active = TRUE").
		Where("deleted_at IS NULL").
		Where("id = ?", productId).
		Limit(1).
		Find(&resp).Error
	return resp, err
}
