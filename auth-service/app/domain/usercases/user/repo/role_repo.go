package repo

import (
	"auth-service/app/domain/entities"
	"auth-service/app/domain/usercases/user/orm"
	"auth-service/app/infra/db"
	"auth-service/app/proto-gen/message"
	"auth-service/app/proto-gen/rpc"
	"auth-service/app/uerror"
	"context"
)

type roleRepo struct {
	rpc.UnimplementedRoleServiceServer
}

var RoleRepo rpc.RoleServiceServer

func init() {
	RoleRepo = &roleRepo{
		UnimplementedRoleServiceServer: rpc.UnimplementedRoleServiceServer{},
	}
}

func (r roleRepo) GetRoleById(ctx context.Context, req *message.RoleRequest) (*message.RoleResponse, error) {
	if req == nil {
		return &message.RoleResponse{}, nil
	}
	roleResp, err := orm.Role.GetRoleById(int(req.Id))
	if err != nil {
		return &message.RoleResponse{}, uerror.InternalError(err, err.Error())
	}
	return r.bind(roleResp), nil
}

func (r roleRepo) CreateRole(ctx context.Context, req *message.RoleRequest) (*message.RoleResponse, error) {
	if req == nil {
		return &message.RoleResponse{}, nil
	}

	tx := db.BeginTx()
	defer db.RecoveryTx(tx)

	roleEntity := &entities.Role{
		Name: req.RoleName,
	}

	err := orm.Role.CreateRoleTx(roleEntity, tx)
	if err != nil {
		tx.Rollback()
		return &message.RoleResponse{}, uerror.InternalError(err, err.Error())
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return &message.RoleResponse{}, uerror.InternalError(err, err.Error())
	}
	return r.bind(roleEntity), nil
}

func (r roleRepo) bind(roleEntity *entities.Role) *message.RoleResponse {
	if roleEntity == nil {
		return &message.RoleResponse{}
	}
	return &message.RoleResponse{
		Id:       int64(roleEntity.Id),
		RoleName: roleEntity.Name,
	}
}
