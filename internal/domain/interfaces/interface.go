package interfaces

import "hokapi/internal/domain/dto/response"

type ServiceInterface interface {
	GetAllHero(search string) ([]response.ResponseHero, error)
	GetHeroById(id string) (response.ResponseHero, error)
}

type RepositoryInterface interface {
	GetAllHero(search string) ([]response.ResponseHero, error)
	GetHeroById(id string) (response.ResponseHero, error)
}