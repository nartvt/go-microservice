package middleware

import (
	"auth-service/app/domain/usercases/repo"
	repo2 "auth-service/app/domain/usercases/user/repo"
)

func GetCurrentUser(userName string) (*repo.UserRepo, error) {
	return repo2.User.GetUserByUserName(userName)
}

func IsAdmin(roleId int) bool {
	return roleId == 1 || roleId == 2
}
