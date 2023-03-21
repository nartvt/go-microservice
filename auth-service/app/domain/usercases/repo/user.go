package repo

import (
	"time"

	"gorm.io/gorm"

	"auth-service/app/domain/entities"
	"auth-service/app/domain/usercases/orm"
	"auth-service/app/uerror"
)

type UserRepo struct {
	Id          int
	UserName    string
	Password    string
	Email       string
	PhoneNumber string
	FullName    string
	Roles       []int
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
type IUserRepo interface {
	GetUserByUserName(userName string) (*UserRepo, error)
}

type user struct{}

var User IUserRepo

func init() {
	User = user{}
}

func (u user) GetUserByUserName(userName string) (*UserRepo, error) {
	userOrm, err := orm.User.GetUserByUserName(userName)
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, uerror.InternalError(err, err.Error())
	}
	if userOrm == nil {
		return nil, nil
	}
	return u.bind(userOrm)
}

func (user) bind(userEntity *entities.User) (*UserRepo, error) {
	if userEntity == nil {
		return nil, nil
	}
	resp := &UserRepo{
		Id:          userEntity.Id,
		UserName:    userEntity.UserName,
		FullName:    userEntity.FullName,
		Email:       userEntity.Email,
		PhoneNumber: userEntity.PhoneNumber,
		CreatedAt:   userEntity.CreatedAt,
		UpdatedAt:   userEntity.UpdatedAt,
	}
	roles, err := orm.Role.GetRoleByUserId(userEntity.Id)
	if err != nil {
		return resp, uerror.InternalError(err, err.Error())
	}

	roleIds := make([]int, len(roles))
	for i := range roles {
		roleIds[i] = roles[i].Id
	}
	resp.Roles = roleIds
	return resp, nil
}
