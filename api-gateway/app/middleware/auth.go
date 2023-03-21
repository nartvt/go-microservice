package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"api-gateway/app/domain/usercases/user/repo"
)

func GetCurrentUser(ctx *fiber.Ctx) (*repo.UserRepo, error) {
	return repo.User.GetUserByUserName(ctx.Get(userNameKey))
}

func IsAdmin(ctx *fiber.Ctx) bool {
	role := ctx.Get(roleKey)
	roleId, _ := strconv.Atoi(role)
	return roleId == 1 || roleId == 2
}
