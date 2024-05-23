package repository

import (
	"errors"

	"hokapi/internal/domain/dto/response"
	"hokapi/internal/domain/interfaces"
	"hokapi/internal/domain/model"

	"github.com/teranixbq/gfunc"
)

type repositoryHero struct {
	repository *gfunc.Query
}

func NewRepositoryHOK(repository *gfunc.Query) interfaces.RepositoryInterface {
	return &repositoryHero{
		repository: repository,
	}
}

func (hok *repositoryHero) GetAllHero(search string) ([]response.ResponseHero, error) {
	data := []model.Hero{}
	var datalist []interface{}

	for _, v := range data {
		datalist = append(datalist, v)
	}

	if search != "" {
		dataSearch, err := hok.repository.FindAllBy(search, "name", datalist)
		if err != nil {
			return nil, err
		}

		if dataSearch == nil {
			return nil, errors.New("error : data not found")
		}
		responseData := response.ListInterfaceToResponse(dataSearch)
		return responseData, nil
	}

	err := hok.repository.Find(&data)
	if err != nil {
		return nil, err
	}

	responseData := response.ListModelToResponse(data)
	return responseData, nil
}

func (hok *repositoryHero) GetHeroById(id string) (response.ResponseHero, error) {
	data := []model.Hero{}
	var datalist []interface{}
	for _, v := range data {
		datalist = append(datalist, v)
	}

	dataById, err := hok.repository.FindBy(id, "id", datalist)
	if err != nil {
		return response.ResponseHero{}, err
	}

	if dataById == nil {
		return response.ResponseHero{}, errors.New("error : data not found")
	}

	resultData := response.InterfaceToResponse(dataById)
	return resultData, nil
}
