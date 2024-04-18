package services

import (
	"AlexSilverson/lab2/src/City/domain/entity"
	"AlexSilverson/lab2/src/City/domain/persistance"
)

type cityService struct {
	cityRepository persistance.CityRepository
}

type CityService interface {
	GetCityById(id uint) *entity.City
	AddCity(city entity.City) *entity.City
	UpdateCity(inputCity entity.City) *entity.City
	DeleteCity(id uint) *entity.City
}

func (service cityService) GetCityById(id uint) *entity.City {

	return service.cityRepository.GetCityById(id)
}

func (service cityService) AddCity(city entity.City) *entity.City {

	return service.cityRepository.AddCity(city)
}

func (service cityService) UpdateCity(inputCity entity.City) *entity.City {

	return service.cityRepository.UpdateCity(inputCity)
}
func (service cityService) DeleteCity(id uint) *entity.City {

	return service.cityRepository.DeleteCity(id)
}

func NewCityRepository(cityRepository persistance.CityRepository) CityService {
	return &cityService{cityRepository: cityRepository}
}
