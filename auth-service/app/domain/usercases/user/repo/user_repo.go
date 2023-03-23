package repo

import (
	"auth-service/app/domain/entities"
	"auth-service/app/domain/usercases/user/orm"
	"auth-service/app/infra/db"
	"auth-service/app/proto-gen/message"
	"auth-service/app/proto-gen/rpc"
	"auth-service/app/uerror"
	"context"
	"gorm.io/gorm"
	"time"
)

type userRepo struct {
	rpc.UnimplementedUserServiceServer
}

var User rpc.UserServiceServer

func init() {
	User = &userRepo{
		UnimplementedUserServiceServer: rpc.UnimplementedUserServiceServer{},
	}
}

func (u userRepo) CreateUser(ctx context.Context, request *message.UserRequest) (*message.UserResponse, error) {
	if request == nil {
		return &message.UserResponse{}, nil
	}
	tx := db.BeginTx()
	defer db.RecoveryTx(tx)

	newUser := entities.User{
		UserName:    request.UserName,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}
	if request.CreatedAt > 0 {
		createdAt := time.UnixMilli(request.CreatedAt)
		newUser.CreatedAt = &createdAt
	}
	err := orm.User.UpdateUserTx(&newUser, tx)
	if err != nil {
		tx.Rollback()
		return &message.UserResponse{}, uerror.InternalError(err, err.Error())
	}
	if request.Role > 0 {
		userRole := &entities.UserRole{
			RoleId: int(request.Role),
			UserId: newUser.Id,
		}
		err = orm.UserRole.CreateUserRoleTx(userRole, tx)
		if err != nil {
			tx.Rollback()
			return &message.UserResponse{}, uerror.InternalError(err, err.Error())
		}
	}
	return u.Bind(&newUser), nil
}

func (u userRepo) UpdateUserInfo(ctx context.Context, request *message.UserRequest) (*message.UserResponse, error) {
	if request == nil {
		return nil, nil
	}
	tx := db.BeginTx()
	defer db.RecoveryTx(tx)

	userOrm, err := orm.User.GetUserByUserNameTx(request.UserName, tx)
	if err != nil && err == gorm.ErrRecordNotFound {
		tx.Rollback()
		return &message.UserResponse{}, err
	}
	if err != nil {
		tx.Rollback()
		return &message.UserResponse{}, uerror.InternalError(err, err.Error())
	}

	if userOrm == nil {
		tx.Rollback()
		return &message.UserResponse{}, err
	}
	if len(request.FirstName) > 0 {
		userOrm.FirstName = request.FirstName
	}
	if len(request.LastName) > 0 {
		userOrm.LastName = request.LastName
	}
	if len(request.PhoneNumber) > 0 {
		userOrm.PhoneNumber = request.PhoneNumber
	}
	now := time.Now()
	if request.UpdatedAt > 0 {
		userOrm.UpdatedAt = &now
	}

	err = orm.User.UpdateUserTx(userOrm, tx)
	if err != nil {
		tx.Rollback()
		return &message.UserResponse{}, uerror.InternalError(err, err.Error())
	}

	if request.Role > 0 {
		err = orm.Role.UpdateUserRoleTx(userOrm.Id, int(request.Role), tx)
		if err != nil {
			tx.Rollback()
			return &message.UserResponse{}, uerror.InternalError(err, err.Error())
		}
	}
	return u.Bind(userOrm), nil
}

func (u userRepo) GetUserByUserName(ctx context.Context, request *message.UserRequest) (*message.UserResponse, error) {
	if request == nil {
		return nil, nil
	}

	userOrm, err := orm.User.GetUserByUserName(request.UserName)
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, uerror.InternalError(err, err.Error())
	}

	if userOrm == nil {
		return nil, nil
	}
	return u.Bind(userOrm), nil
}
func (u userRepo) Ping(ctx context.Context, request *message.UserRequest) (*message.UserResponse, error) {
	return &message.UserResponse{Name: "ping response from server"}, nil
}

func (userRepo) Bind(userEntity *entities.User) *message.UserResponse {
	if userEntity == nil {
		return nil
	}
	return &message.UserResponse{
		Id:          int64(userEntity.Id),
		UserName:    userEntity.UserName,
		FullName:    userEntity.FullName,
		Email:       userEntity.Email,
		PhoneNumber: userEntity.PhoneNumber,
		CreatedAt:   userEntity.CreatedAt.UnixMilli(),
		UpdatedAt:   userEntity.UpdatedAt.UnixMilli(),
		Role:        1,
	}
}
