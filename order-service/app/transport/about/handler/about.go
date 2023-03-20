package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"order-service/app/domain/usercases/about/repo"
	"order-service/app/transport"
	"order-service/app/transport/about/request"
	"order-service/app/transport/about/response"
	"order-service/app/uerror"
)

type AboutHandler struct {
	AboutDomain repo.IAboutRepo
}

func (s AboutHandler) GetAboutBySectionId(ctx *fiber.Ctx) error {
	param := &request.AboutRequestParam{}
	if err := param.Bind(ctx); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(err)
	}
	abouts, total, err := s.AboutDomain.GetAboutBySectionId(param.SectionId, param.Limit, param.Offset)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(err)
	}

	if len(abouts) <= 0 {
		return ctx.Status(http.StatusNotFound).
			JSON(uerror.NotFoundError(fmt.Errorf("not found %s", "section"), "section not found"))
	}

	resp := transport.Response{
		Data: response.NewAboutResponses(abouts),
	}
	if len(abouts) < param.Limit {
		return ctx.Status(http.StatusOK).JSON(resp)
	}
	nextUrl := fmt.Sprintf("%s?limit=%d&page=%d", ctx.OriginalURL(), param.Limit, param.Page+1)
	resp.Pagination = &transport.Pagination{
		NextUrl: nextUrl,
		Total:   total,
	}
	return ctx.Status(http.StatusOK).JSON(resp)
}
