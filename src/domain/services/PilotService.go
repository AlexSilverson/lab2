package services

import (
	"AlexSilverson/lab2/src/domain/entity"

	"gorm.io/gorm"
)

type pilotService struct {
	db *gorm.DB
}

type PilotService interface {
	GetPilotById(id uint) (*entity.Pilot, error)
	AddPilot(pilot entity.Pilot) error
	UpdatePilot(inputPilot entity.Pilot) error
	DeletePilot(id uint) error
}

func (r pilotService) GetPilotById(id uint) (*entity.Pilot, error) {
	var pilot entity.Pilot

	erdb := r.db.First(&pilot, id)
	if erdb.Error != nil {
		return &pilot, erdb.Error
	}
	return &pilot, nil
}

func (r pilotService) AddPilot(pilot entity.Pilot) error {
	erdb := r.db.Create(&pilot)
	if erdb.Error != nil {
		return erdb.Error
	}
	return nil
}

func (r pilotService) UpdatePilot(inputPilot entity.Pilot) error {
	erdb := r.db.Save(&inputPilot)
	if erdb.Error != nil {
		return erdb.Error
	}
	return nil
}

func (r pilotService) DeletePilot(id uint) error {
	erdb := r.db.Delete(&entity.Pilot{}, id)
	if erdb.Error != nil {
		return erdb.Error
	}
	return nil

}

func NewPilotSevice(db *gorm.DB) PilotService {
	return &pilotService{db: db}
}
