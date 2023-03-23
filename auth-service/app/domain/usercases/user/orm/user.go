package orm

import (
	"auth-service/app/domain/entities"
	"auth-service/app/domain/interfaces"
	"auth-service/app/infra/db"
	"gorm.io/gorm"
)

type user struct{}

var User interfaces.IUser

func init() {
	User = user{}
}

func (u user) CreateUserTx(userCreate *entities.User, tx *gorm.DB) error {
	return tx.Create(userCreate).Error
}

func (u user) UpdateUserTx(userUpdate *entities.User, tx *gorm.DB) error {
	return tx.Save(userUpdate).Error
}

func (u user) GetUserById(id int) (*entities.User, error) {
	resp := &entities.User{}
	err := db.DB().Model(&entities.User{}).
		Where("id = ?", id).
		Limit(1).
		Find(resp).Error
	return resp, err
}

func (u user) GetUserByEmail(email string) (*entities.User, error) {
	resp := &entities.User{}
	err := db.DB().Model(&entities.User{}).
		Where("email = ?", email).
		Limit(1).
		Find(resp).Error
	return resp, err
}

func (u user) GetUserByUserNameTx(userName string, tx *gorm.DB) (*entities.User, error) {
	resp := &entities.User{}
	err := tx.Model(&entities.User{}).
		Where("user_name = ?", userName).
		Limit(1).
		Find(resp).Error
	return resp, err
}

func (u user) GetUserByUserName(userName string) (*entities.User, error) {
	resp := &entities.User{}
	err := db.DB().Model(&entities.User{}).
		Where("user_name = ?", userName).
		Limit(1).
		Find(resp).Error
	return resp, err
}

func (u user) GetUserByPhone(phoneNumber string) (*entities.User, error) {
	resp := &entities.User{}
	err := db.DB().Model(&entities.User{}).
		Where("phone_number = ?", phoneNumber).
		Limit(1).
		Find(resp).Error
	return resp, err
}
