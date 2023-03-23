package middleware

import (
	"api-gateway/app/domain/models"
	"api-gateway/app/domain/usercases/user/repo"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetCurrentUser(userName string) (*models.UserRepo, error) {
	return repo.User.GetUserByUserName(userName)
}

func IsAdmin(ctx *fiber.Ctx) bool {
	role := ctx.Get(roleKey)
	roleId, _ := strconv.Atoi(role)
	return roleId == 1 || roleId == 2
}
