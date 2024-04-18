package services

import (
	"AlexSilverson/lab2/src/City/domain/entity"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type cityService struct {
	fileRoute string
}

type CityService interface {
	GetCityById(id uint) (*entity.City, error)
	AddCity(city entity.City) error
	UpdateCity(inputCity entity.City) error
	DeleteCity(id uint) error
}

func (r cityService) GetCityById(id uint) (*entity.City, error) {
	var city entity.City

	f, er := os.Open(r.fileRoute)

	if er != nil {
		return &city, er
	}

	defer f.Close()

	decoder := json.NewDecoder(f)

	for decoder.More() {
		er = decoder.Decode(&city)
		if er != nil {
			return &city, er
		}
		if city.ID == id {
			return &city, nil
		}
	}
	return &city, errors.New("city not found")
}

func (r cityService) AddCity(city entity.City) error {
	f, er := os.OpenFile(r.fileRoute, os.O_WRONLY|os.O_APPEND, 0644)

	if er != nil {
		return er
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	data, er := json.Marshal(city)

	if er != nil {
		return er
	}

	_, err := w.Write(data)

	if err != nil {
		return er
	}
	er = w.Flush()
	if er != nil {
		return er
	}
	return nil
}

func (r cityService) UpdateCity(inputCity entity.City) error {
	f, er := os.OpenFile(r.fileRoute, os.O_APPEND|os.O_RDWR, 0644)

	if er != nil {
		return er
	}

	defer f.Close()

	citis := make([]entity.City, 0)

	decoder := json.NewDecoder(f)

	for decoder.More() {
		var now entity.City
		er = decoder.Decode(&now)
		if er != nil {
			return er
		}
		citis = append(citis, now)
	}

	er = os.Truncate(r.fileRoute, 0)
	if er != nil {
		return er
	}
	var flag bool = false
	fmt.Println(citis)
	for _, curCity := range citis {
		if curCity.ID == inputCity.ID {
			flag = true
			r.AddCity(inputCity)

		} else {
			r.AddCity(curCity)
		}
	}

	if flag {
		return nil
	} else {
		return errors.New("that city not found")
	}
}

func (r cityService) DeleteCity(id uint) error {
	f, er := os.OpenFile(r.fileRoute, os.O_APPEND|os.O_RDWR, 0644)

	if er != nil {
		return er
	}

	defer f.Close()

	citis := make([]entity.City, 0)

	decoder := json.NewDecoder(f)

	for decoder.More() {
		var now entity.City
		er = decoder.Decode(&now)
		fmt.Println(now)
		if er != nil {
			return (er)
		}
		citis = append(citis, now)
	}

	er = os.Truncate(r.fileRoute, 0)
	if er != nil {
		return er
	}
	fmt.Println(citis)
	for _, curCity := range citis {
		if curCity.ID != id {
			r.AddCity(curCity)
		}
	}
	return nil

}

func NewCitySevice(fileRoute string) CityService {
	return &cityService{fileRoute: fileRoute}
}
