package repo

import (
	"api-gateway/app/domain/interfaces"
	"api-gateway/app/domain/models"
	"auth-service/app/proto-gen/message"
	client "auth-service/app/transport/grpc"
	"time"
)

type user struct{}

var User interfaces.IUserRepo

func init() {
	User = user{}
}

func (u user) GetUserByUserName(userName string) (*models.UserRepo, error) {
	userMessage, err := client.User.GetUserByUserName(userName)
	if err != nil {
		return nil, err
	}
	if userMessage == nil {
		return nil, nil
	}
	return u.Bind(userMessage), nil
}

func (user) Bind(userMessage *message.UserResponse) *models.UserRepo {
	if userMessage == nil {
		return nil
	}
	resp := &models.UserRepo{
		Id:          int(userMessage.Id),
		UserName:    userMessage.UserName,
		Email:       userMessage.Email,
		PhoneNumber: userMessage.PhoneNumber,
		FullName:    userMessage.FullName,
		Role:        int(userMessage.Role),
	}
	createdAt := time.UnixMilli(userMessage.CreatedAt)
	updatedAt := time.UnixMilli(userMessage.UpdatedAt)
	resp.CreatedAt = &createdAt
	resp.UpdatedAt = &updatedAt
	return resp
}
