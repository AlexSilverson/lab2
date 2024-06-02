package dtos

import "AlexSilverson/lab2/src/domain/entity"

type AddFlightDTO struct {
	From       uint   `json:"from" validate:"required"`
	To         uint   `json:"to" validate:"required"`
	TrevelTime uint   `json:"treveltime" validate:"required"`
	Pilot      uint   `json:"pilot" validate:"required"`
	StartTime  string `json:"starttime" validate:"required"`
}

func (dto AddFlightDTO) MapAddFlightDtoToFlight() entity.Flight {
	var ans entity.Flight
	ans.FromId = dto.From
	ans.ToId = dto.To
	ans.Pilot.ID = dto.Pilot
	ans.StartTime = dto.StartTime
	return ans
}
