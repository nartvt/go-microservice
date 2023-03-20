package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"order-service/app/domain/usercases/user/repo"
	"order-service/app/transport"
	"order-service/app/transport/user/response"
)

type UserHandler struct {
	UserDomain repo.IUserRepo
}

func (u UserHandler) GetMe(ctx *fiber.Ctx) error {
	user := ctx.Locals("user")
	if user == nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("user invalid")})
	}
	userRepo := user.(repo.UserRepo)
	return ctx.Status(http.StatusOK).JSON(transport.Response{
		Data: response.NewUserResponse(userRepo),
	})
}
