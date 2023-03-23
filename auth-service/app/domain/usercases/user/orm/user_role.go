package orm

import (
	"auth-service/app/domain/entities"
	"auth-service/app/domain/interfaces"
	"gorm.io/gorm"
)

type userRoleOrm struct{}

var UserRole interfaces.IUserRole

func init() {
	UserRole = userRoleOrm{}
}

func (u userRoleOrm) CreateUserRoleTx(userRoleEntity *entities.UserRole, tx *gorm.DB) error {
	return tx.Create(userRoleEntity).Error
}

func (u userRoleOrm) UpdateUserRoleTx(userRoleEntity *entities.UserRole, tx *gorm.DB) error {
	return tx.Save(userRoleEntity).Error
}
