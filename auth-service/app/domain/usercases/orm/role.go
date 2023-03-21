package orm

import (
	"auth-service/app/domain/entities"
	"auth-service/app/infra/db"
)

type IRole interface {
	GetRoleByUserId(userId int) ([]entities.Role, error)
}
type role struct{}

var Role IRole

func init() {
	Role = role{}
}

func (role) GetRoleByUserId(userId int) ([]entities.Role, error) {
	var resp []entities.Role
	err := db.DB().Model(&entities.Role{}).
		Joins("JOIN user_roles ON roles.id = user_roles.role_id").
		Joins("JOIN users ON users.id = user_roles.user_id").
		Where("user_id = ?", userId).
		Find(resp).Error
	return resp, err
}
