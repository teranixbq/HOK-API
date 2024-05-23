package route

import (
	"hokapi/domain/handler"
	"hokapi/domain/repository"
	"hokapi/domain/service"

	"github.com/gofiber/fiber/v2"
	"github.com/teranixbq/gfunc"
)

func InitRoute(f *fiber.App, g *gfunc.Query) {

	repositoryHOK := repository.NewRepositoryHOK(g)
	serviceHOK := service.NewRepositoryHOK(repositoryHOK)
	handlerHOK := handler.NewHandlerHOK(serviceHOK)

	f.Get("/", handlerHOK.GetAllHero)
	f.Get("/:id", handlerHOK.GetHeroById)
}
