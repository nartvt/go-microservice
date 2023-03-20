package orm

import (
	"order-service/app/domain/entities"
	"order-service/app/infra/db"
)

type IUserBodyRecord interface {
	GetBodyRecordByUserId(userId int, limit int, offset int) ([]entities.UserBodyRecord, error)
}
type userBodyRecord struct{}

var UserBodyRecord IUserBodyRecord

func init() {
	UserBodyRecord = userBodyRecord{}
}
func (u userBodyRecord) GetBodyRecordByUserId(userId int, limit int, offset int) ([]entities.UserBodyRecord, error) {
	resp := []entities.UserBodyRecord{}
	query := db.DB().Model(&entities.UserBodyRecord{}).
		Limit(limit).
		Offset(offset).Order("id DESC")
	if userId > 0 {
		query = query.Where("user_id = ?", userId)
	}
	err := query.Find(&resp).Error
	return resp, err
}
