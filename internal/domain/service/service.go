package service

import (
	"errors"
	"strconv"

	"hokapi/internal/domain/dto/response"
	"hokapi/internal/domain/interfaces"
)

type serviceHero struct {
	service interfaces.RepositoryInterface
}

func NewRepositoryHOK(service interfaces.RepositoryInterface) interfaces.ServiceInterface {
	return &serviceHero{
		service: service,
	}
}

func (hok *serviceHero) GetAllHero(search string) ([]response.ResponseHero, error) {
	result, err := hok.service.GetAllHero(search)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (hok *serviceHero) GetHeroById(id string) (response.ResponseHero, error) {
	_, errId := strconv.ParseUint(id, 10, 64)
	if errId != nil {
		return response.ResponseHero{}, errors.New("error : id must be a number")
	}

	result, err := hok.service.GetHeroById(id)
	if err != nil {
		return response.ResponseHero{}, err
	}

	return result, nil
}
