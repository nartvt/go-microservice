package interfaces

import (
	"auth-service/app/domain/entities"
	"gorm.io/gorm"
)

type IUser interface {
	GetUserById(id int) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByPhone(phoneNumber string) (*entities.User, error)
	GetUserByUserNameTx(userName string, tx *gorm.DB) (*entities.User, error)
	GetUserByUserName(userName string) (*entities.User, error)
	UpdateUserTx(userUpdate *entities.User, tx *gorm.DB) error
	CreateUserTx(userCreate *entities.User, tx *gorm.DB) error
}
