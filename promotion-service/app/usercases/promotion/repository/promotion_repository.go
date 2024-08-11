package repository

import (
	"gorm.io/gorm"
	"promotion-service/app/domain/model"
	"promotion-service/app/infra/db"
)

type promotionRepository struct {
}

type IPromotionRepository interface {
	GetPromotionById(promotionId int, isActive bool) (*model.Promotion, error)
	GetPromotionsPagination(isActive bool, limit int, offset int) ([]model.Promotion, error)
	UpdatePromotionTx(promotion *model.Promotion, tx *gorm.DB) error
	CreatePromotionTx(promotion *model.Promotion, tx *gorm.DB) error
}

var PromotionRepository IPromotionRepository

func init() {
	PromotionRepository = &promotionRepository{}
}

func (p promotionRepository) UpdatePromotionTx(promotion *model.Promotion, tx *gorm.DB) error {
	return tx.Save(promotion).Error
}

func (p promotionRepository) CreatePromotionTx(promotion *model.Promotion, tx *gorm.DB) error {
	return tx.Create(promotion).Error
}

func (p promotionRepository) GetPromotionById(promotionId int, isActive bool) (*model.Promotion, error) {
	if promotionId <= 0 {
		return nil, nil
	}
	var resp *model.Promotion
	err := db.DB().Model(&model.Promotion{}).
		Where("active = ?", isActive).
		Where("id = ?", promotionId).
		Where("deleted_at IS NULL").
		Find(resp).
		Error
	return resp, err
}
func (p promotionRepository) GetPromotionsPagination(isActive bool, limit int, offset int) ([]model.Promotion, error) {
	var resp []model.Promotion
	err := db.DB().Model(&model.Promotion{}).Where("deleted_at IS NULL").
		Where("active = ?", isActive).
		Order("id DESC").
		Offset(offset).
		Limit(limit).
		Find(&resp).
		Error
	return resp, err
}
