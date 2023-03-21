package router

import (
	"api-gateway/app/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	v1.Use(middleware.CorsFilter(), middleware.RateLimit())

	setupUserRoute(v1)
	setupNewsfeedRoute(v1.Group("/user/me", middleware.Auth.RequireLogin()))
	setupAboutRoute(v1)
}
