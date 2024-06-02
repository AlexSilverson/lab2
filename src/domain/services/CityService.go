package services

import (
	"AlexSilverson/lab2/src/domain/entity"

	"gorm.io/gorm"
)

type cityService struct {
	db *gorm.DB
}

type CityService interface {
	GetCityById(id uint) (*entity.City, error)
	AddCity(city entity.City) error
	UpdateCity(inputCity entity.City) error
	DeleteCity(id uint) error
}

func (r cityService) GetCityById(id uint) (*entity.City, error) {
	var city entity.City
	erdb := r.db.First(&city, id)
	if erdb.Error != nil {
		return &city, erdb.Error
	}
	return &city, nil
}

func (r cityService) AddCity(city entity.City) error {
	erdb := r.db.Create(&city)
	if erdb.Error != nil {
		return erdb.Error
	}
	return nil
}

func (r cityService) UpdateCity(inputCity entity.City) error {
	erdb := r.db.Save(&inputCity)
	if erdb.Error != nil {
		return erdb.Error
	}
	return nil
}

func (r cityService) DeleteCity(id uint) error {
	erdb := r.db.Delete(&entity.City{}, id)
	if erdb.Error != nil {
		return erdb.Error
	}
	return nil
}

func NewCitySevice(db *gorm.DB) CityService {
	return &cityService{db: db}
}
