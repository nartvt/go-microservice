package interfaces

import "api-gateway/app/domain/models"

type IUserRepo interface {
	GetUserByUserName(userName string) (*models.UserRepo, error)
}
