package orm

import (
	"api-gateway/app/domain/entities"
	"api-gateway/app/infra/db"
)

type IUserDiary interface {
	GetUserDiaryByUserId(userId int, limit int, offset int) ([]entities.UserDiary, error)
}
type userDiary struct{}

var UserDiary IUserDiary

func init() {
	UserDiary = userDiary{}
}
func (u userDiary) GetUserDiaryByUserId(userId int, limit int, offset int) ([]entities.UserDiary, error) {
	if userId <= 0 {
		return []entities.UserDiary{}, nil
	}
	resp := []entities.UserDiary{}
	err := db.DB().Model(&entities.UserDiary{}).
		Where("user_id = ?", userId).
		Limit(limit).
		Offset(offset).Order("id DESC").
		Find(&resp).Error
	return resp, err
}
