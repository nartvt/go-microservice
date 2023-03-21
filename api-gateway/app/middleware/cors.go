package middleware

import (
	"api-gateway/app/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"time"
)

func CorsFilter() fiber.Handler {
	return cors.New(*corsConfig())
}

const appName = "api gateway version 1.0"

func corsConfig() *cors.Config {
	return &cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,HEAD,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}
}

func Config() fiber.Config {
	return fiber.Config{
		AppName:           appName,
		EnablePrintRoutes: true,
		ReadTimeout:       time.Duration(config.Get().Server.ReadTimeOut) * time.Millisecond,
		WriteTimeout:      time.Duration(config.Get().Server.WriteTimeOut) * time.Millisecond,
	}
}
