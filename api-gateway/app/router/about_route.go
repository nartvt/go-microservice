package router

import (
	aboutRepo "api-gateway/app/domain/usercases/about/repo"
	aboutHandler "api-gateway/app/transport/about/handler"

	"github.com/gofiber/fiber/v2"
)

func setupAboutRoute(v1 fiber.Router) {
	groupAboutHandler := aboutHandler.AboutHandler{
		AboutDomain: aboutRepo.NewAboutRepo(),
	}

	groupAbout := v1.Group("/abouts")
	{
		GET(groupAbout, "/:sectionId", groupAboutHandler.GetAboutBySectionId)
	}
}
