package services

import (
	"AlexSilverson/lab2/src/domain/entity"
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type flightService struct {
	fileRoute string
}

type FlightService interface {
	GetFlightById(id uint) (*entity.Flight, error)
	AddFlight(flight entity.Flight) error
	UpdateFlight(inputFlight entity.Flight) error
	DeleteFlight(id uint) error
}

func NewFlightSevice(fileRoute string) FlightService {
	return &flightService{fileRoute: fileRoute}
}

func (r flightService) GetFlightById(id uint) (*entity.Flight, error) {
	var flight entity.Flight

	f, er := os.Open(r.fileRoute)

	if er != nil {
		return &flight, er
	}

	defer f.Close()

	decoder := json.NewDecoder(f)

	for decoder.More() {
		er = decoder.Decode(&flight)
		if er != nil {
			return &flight, er
		}
		if flight.ID == id {
			return &flight, nil
		}
	}
	return &flight, errors.New("flight not found")
}

func (r flightService) AddFlight(flight entity.Flight) error {
	f, er := os.OpenFile(r.fileRoute, os.O_WRONLY|os.O_APPEND, 0644)

	if er != nil {
		return er
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	data, er := json.Marshal(flight)

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

func (r flightService) UpdateFlight(inputFlight entity.Flight) error {
	f, er := os.OpenFile(r.fileRoute, os.O_APPEND|os.O_RDWR, 0644)

	if er != nil {
		return er
	}

	defer f.Close()

	flights := make([]entity.Flight, 0)

	decoder := json.NewDecoder(f)

	for decoder.More() {
		var now entity.Flight
		er = decoder.Decode(&now)
		if er != nil {
			return er
		}
		flights = append(flights, now)
	}

	er = os.Truncate(r.fileRoute, 0)
	if er != nil {
		return er
	}
	var flag bool = false
	for _, curFlight := range flights {
		if curFlight.ID == inputFlight.ID {
			flag = true
			r.AddFlight(inputFlight)

		} else {
			r.AddFlight(curFlight)
		}
	}

	if flag {
		return nil
	} else {
		return errors.New("that flight not found")
	}
}

func (r flightService) DeleteFlight(id uint) error {
	f, er := os.OpenFile(r.fileRoute, os.O_APPEND|os.O_RDWR, 0644)

	if er != nil {
		return er
	}

	defer f.Close()

	flights := make([]entity.Flight, 0)

	decoder := json.NewDecoder(f)

	for decoder.More() {
		var now entity.Flight
		er = decoder.Decode(&now)
		if er != nil {
			return (er)
		}
		flights = append(flights, now)
	}

	er = os.Truncate(r.fileRoute, 0)
	if er != nil {
		return er
	}
	for _, curFlight := range flights {
		if curFlight.ID != id {
			r.AddFlight(curFlight)
		}
	}
	return nil

}
