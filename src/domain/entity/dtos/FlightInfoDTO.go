package dtos

import "AlexSilverson/lab2/src/domain/entity"

type FlightInfoDto struct {
	From       string `json:"from"`
	To         string `json:"to"`
	Treveltime uint   `json:"treveltime"`
	Starttime  string `json:"starttime"`
}

func MapToFlightInfoDto(flight entity.Flight) FlightInfoDto {
	var dto FlightInfoDto
	dto.From = flight.From.Name
	dto.To = flight.To.Name
	dto.Treveltime = flight.TrevelTime
	dto.Starttime = flight.StartTime
	return dto
}
