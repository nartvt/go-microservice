package interfaces

import (
	"auth-service/app/domain/entities"
	"gorm.io/gorm"
)

type IUserRole interface {
	CreateUserRoleTx(userRoleEntity *entities.UserRole, tx *gorm.DB) error
	UpdateUserRoleTx(userRoleEntity *entities.UserRole, tx *gorm.DB) error
}

type IRole interface {
	GetRoleByUserId(userId int) ([]entities.Role, error)
	GetRoleById(roleId int) (*entities.Role, error)
	CreateRoleTx(role *entities.Role, tx *gorm.DB) error
	UpdateUserRoleTx(userId int, roleId int, tx *gorm.DB) error
}
