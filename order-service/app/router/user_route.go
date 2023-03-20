package router

import (
	userrepo "order-service/app/domain/usercases/user/repo"
	"order-service/app/middleware"
	userHandler "order-service/app/transport/user/handler"

	"github.com/gofiber/fiber/v2"
)

func setupUserRoute(v1 fiber.Router) {
	groupUserHandler := userHandler.UserHandler{
		UserDomain: userrepo.User,
	}

	groupUser := v1.Group("/user/me", middleware.Auth.RequireLogin())
	{
		GET(groupUser, "", groupUserHandler.GetMe)
	}
}
