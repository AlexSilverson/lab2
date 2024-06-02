package services

import (
	"AlexSilverson/lab2/src/domain/entity"

	"gorm.io/gorm"
)

type flightService struct {
	db *gorm.DB
}

type FlightService interface {
	GetFlightById(id uint) (*entity.Flight, error)
	AddFlight(flight entity.Flight) error
	UpdateFlight(inputFlight entity.Flight) error
	DeleteFlight(id uint) error
}

func NewFlightSevice(db *gorm.DB) FlightService {
	return &flightService{db: db}
}

func (r flightService) GetFlightById(id uint) (*entity.Flight, error) {
	var flight entity.Flight
	erdb := r.db.First(&flight, id)
	if erdb.Error != nil {
		return &flight, erdb.Error
	}
	erdb = r.db.First(&flight.To, flight.ToId)
	if erdb.Error != nil {
		return &flight, erdb.Error
	}
	erdb = r.db.First(&flight.From, flight.FromId)
	if erdb.Error != nil {
		return &flight, erdb.Error
	}
	erdb = r.db.First(&flight.Pilot, flight.PilotId)
	if erdb.Error != nil {
		return &flight, erdb.Error
	}
	return &flight, nil
}

func (r flightService) AddFlight(flight entity.Flight) error {
	erdb := r.db.Create(&flight)
	if erdb.Error != nil {
		return erdb.Error
	}
	return nil
}

func (r flightService) UpdateFlight(inputFlight entity.Flight) error {
	erdb := r.db.Save(&inputFlight)
	if erdb.Error != nil {
		return erdb.Error
	}
	return nil
}

func (r flightService) DeleteFlight(id uint) error {
	erdb := r.db.Delete(&entity.Flight{}, id)
	if erdb.Error != nil {
		return erdb.Error
	}
	return nil

}
