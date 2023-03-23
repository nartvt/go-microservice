package orm

import (
	"auth-service/app/domain/entities"
	"auth-service/app/domain/interfaces"
	"auth-service/app/infra/db"
	"gorm.io/gorm"
)

type roleOrm struct{}

var Role interfaces.IRole

func init() {
	Role = roleOrm{}
}

func (roleOrm) CreateRoleTx(role *entities.Role, tx *gorm.DB) error {
	return tx.Create(role).Error

}

func (roleOrm) UpdateUserRoleTx(userId int, roleId int, tx *gorm.DB) error {
	return tx.Model(&entities.UserRole{}).
		Set("role_id = ?", roleId).
		Where("user_id = ?", userId).Error
}

func (roleOrm) GetRoleById(roleId int) (*entities.Role, error) {
	resp := &entities.Role{}
	err := db.DB().Model(&entities.Role{}).
		Where("id = ?", roleId).
		Find(resp).Error
	return resp, err
}

func (roleOrm) GetRoleByUserId(userId int) ([]entities.Role, error) {
	var resp []entities.Role
	err := db.DB().Model(&entities.Role{}).
		Joins("JOIN user_roles ON roles.id = user_roles.role_id").
		Joins("JOIN users ON users.id = user_roles.user_id").
		Where("user_id = ?", userId).
		Find(resp).Error
	return resp, err
}
