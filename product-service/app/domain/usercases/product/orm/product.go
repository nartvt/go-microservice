package orm

import (
	"gorm.io/gorm"
	"product-service/app/domain/entities"
	"product-service/app/domain/interfaces"
	"product-service/app/infra/db"
)

type product struct{}

var Product interfaces.IProduct

func init() {
	Product = product{}
}
func (d product) CreateProductTx(newProduct *entities.Product, tx *gorm.DB) error {
	return tx.Save(newProduct).Error
}

func (d product) GetProductById(productId int) (*entities.Product, error) {
	var resp *entities.Product
	err := db.DB().Model(&entities.Product{}).
		Where("product.active = TRUE").
		Where("deleted_at IS NULL").
		Where("id = ?", productId).
		Limit(1).
		Find(&resp).Error
	return resp, err
}
