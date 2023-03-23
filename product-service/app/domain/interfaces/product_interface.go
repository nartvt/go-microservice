package interfaces

import (
	"gorm.io/gorm"
	"product-service/app/domain/entities"
)

type IProduct interface {
	CreateProductTx(newProduct *entities.Product, tx *gorm.DB) error
	GetProductById(productId int) (*entities.Product, error)
}
