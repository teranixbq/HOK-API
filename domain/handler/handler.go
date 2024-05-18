package handler

import (
	"strings"

	"hokapi/domain/interfaces"
	"hokapi/helper"

	"github.com/gofiber/fiber/v2"
)

type handlerHero struct {
	service interfaces.ServiceInterface
}

func NewHandlerHOK(service interfaces.ServiceInterface) *handlerHero {
	return &handlerHero{
		service: service,
	}
}

func (hok *handlerHero) GetAllHero(f *fiber.Ctx) error {
	search := f.Query("search")

	response, err := hok.service.GetAllHero(search)
	if err != nil {
		if strings.Contains(err.Error(), "error : data not found") {
			return f.Status(fiber.StatusNotFound).JSON(helper.ErrorResponse(err.Error()))
		}
		
		if strings.Contains(err.Error(), "error") {
			return f.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(err.Error()))
		}

		return f.Status(fiber.StatusInternalServerError).JSON(helper.ErrorResponse(err.Error()))
	}

	if len(response) == 0 {
		return f.Status(fiber.StatusOK).JSON(helper.SuccessResponse("data not yet available"))
	}

	return f.Status(fiber.StatusOK).JSON(helper.SuccessWithDataResponse("success get all hero", response))
}

func (hok *handlerHero) GetHeroById(f *fiber.Ctx) error {
	id := f.Params("id")

	response, err := hok.service.GetHeroById(id)
	if err != nil {
		if strings.Contains(err.Error(), "error : data not found") {
			return f.Status(fiber.StatusNotFound).JSON(helper.ErrorResponse(err.Error()))
		}

		if strings.Contains(err.Error(), "error") {
			return f.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(err.Error()))
		}

		return f.Status(fiber.StatusInternalServerError).JSON(helper.ErrorResponse(err.Error()))
	}

	return f.Status(fiber.StatusOK).JSON(helper.SuccessWithDataResponse("success get hero", response))
}
