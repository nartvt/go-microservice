package router

import (
	aboutRepo "api-gateway/app/domain/usercases/about/repository"
	aboutHandler "api-gateway/app/transport/about/handler"
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
