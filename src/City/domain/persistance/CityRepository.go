package persistance

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/AlexSilverson/secondLab/src/City/domain/entity"
)

type cityRepository struct {
	fileRoute string
}

type CityRepository interface {
	GetCityById(id uint) *entity.City
	AddCity(city entity.City) *entity.City
	UpdateCity(inputCity entity.City) *entity.City
	DeleteCity(id uint) *entity.City
}

func (r cityRepository) GetCityById(id uint) *entity.City {
	var city entity.City

	f, er := os.Open(r.fileRoute)

	if er != nil {
		panic("file not found")
	}

	defer f.Close()

	decoder := json.NewDecoder(f)

	for decoder.More() {
		er = decoder.Decode(&city)
		if er != nil {
			continue
		}
		if city.ID == id {
			return &city
		}
	}
	panic("city not found")
}

func (r cityRepository) AddCity(city entity.City) *entity.City {
	f, er := os.OpenFile(r.fileRoute, os.O_WRONLY|os.O_APPEND, 0644)

	if er != nil {
		panic("file not found")
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	data, er := json.Marshal(city)

	if er != nil {
		panic("city cant be marshaled")
	}

	_, err := w.Write(data)

	if err != nil {
		panic(" marshaled city cant be written in a file")
	}
	errr := w.Flush()
	if errr != nil {
		panic(errr)
	}
	return &city
}

func (r cityRepository) UpdateCity(inputCity entity.City) *entity.City {
	f, er := os.OpenFile(r.fileRoute, os.O_APPEND|os.O_RDWR, 0644)

	if er != nil {
		panic("file not found")
	}

	defer f.Close()

	citis := make([]entity.City, 0)

	decoder := json.NewDecoder(f)

	for decoder.More() {
		var now entity.City
		er = decoder.Decode(&now)
		fmt.Println(now)
		if er != nil {
			panic(er)
		}
		citis = append(citis, now)
	}

	er = os.Truncate(r.fileRoute, 0)
	if er != nil {
		panic(er)
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
		return &inputCity
	} else {
		panic("that city wasnt found")
	}
}

func (r cityRepository) DeleteCity(id uint) *entity.City {
	f, er := os.OpenFile(r.fileRoute, os.O_APPEND|os.O_RDWR, 0644)

	if er != nil {
		panic("file not found")
	}

	defer f.Close()

	citis := make([]entity.City, 0)

	decoder := json.NewDecoder(f)

	for decoder.More() {
		var now entity.City
		er = decoder.Decode(&now)
		fmt.Println(now)
		if er != nil {
			panic(er)
		}
		citis = append(citis, now)
	}

	er = os.Truncate(r.fileRoute, 0)
	if er != nil {
		panic(er)
	}
	var flag bool = false
	fmt.Println(citis)
	var delCity entity.City
	for _, curCity := range citis {
		if curCity.ID == id {
			delCity = curCity
			flag = true
		} else {
			r.AddCity(curCity)
		}
	}

	if flag {
		return &delCity
	} else {
		panic("that city wasnt found")
	}
}

func NewCityRepository(fileRoute string) CityRepository {
	return &cityRepository{fileRoute: fileRoute}
}
