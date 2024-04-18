package services

import (
	"AlexSilverson/lab2/src/domain/entity"
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type pilotService struct {
	fileRoute string
}

type PilotService interface {
	GetPilotById(id uint) (*entity.Pilot, error)
	AddPilot(pilot entity.Pilot) error
	UpdatePilot(inputPilot entity.Pilot) error
	DeletePilot(id uint) error
}

func (r pilotService) GetPilotById(id uint) (*entity.Pilot, error) {
	var pilot entity.Pilot

	f, er := os.Open(r.fileRoute)

	if er != nil {
		return &pilot, er
	}

	defer f.Close()

	decoder := json.NewDecoder(f)

	for decoder.More() {
		er = decoder.Decode(&pilot)
		if er != nil {
			return &pilot, er
		}
		if pilot.ID == id {
			return &pilot, nil
		}
	}
	return &pilot, errors.New("pilot not found")
}

func (r pilotService) AddPilot(pilot entity.Pilot) error {
	f, er := os.OpenFile(r.fileRoute, os.O_WRONLY|os.O_APPEND, 0644)

	if er != nil {
		return er
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	data, er := json.Marshal(pilot)

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

func (r pilotService) UpdatePilot(inputPilot entity.Pilot) error {
	f, er := os.OpenFile(r.fileRoute, os.O_APPEND|os.O_RDWR, 0644)

	if er != nil {
		return er
	}

	defer f.Close()

	pilots := make([]entity.Pilot, 0)

	decoder := json.NewDecoder(f)

	for decoder.More() {
		var now entity.Pilot
		er = decoder.Decode(&now)
		if er != nil {
			return er
		}
		pilots = append(pilots, now)
	}

	er = os.Truncate(r.fileRoute, 0)
	if er != nil {
		return er
	}
	var flag bool = false
	for _, curPilot := range pilots {
		if curPilot.ID == inputPilot.ID {
			flag = true
			r.AddPilot(inputPilot)

		} else {
			r.AddPilot(curPilot)
		}
	}

	if flag {
		return nil
	} else {
		return errors.New("that pilot not found")
	}
}

func (r pilotService) DeletePilot(id uint) error {
	f, er := os.OpenFile(r.fileRoute, os.O_APPEND|os.O_RDWR, 0644)

	if er != nil {
		return er
	}

	defer f.Close()

	pilots := make([]entity.Pilot, 0)

	decoder := json.NewDecoder(f)

	for decoder.More() {
		var now entity.Pilot
		er = decoder.Decode(&now)
		if er != nil {
			return (er)
		}
		pilots = append(pilots, now)
	}

	er = os.Truncate(r.fileRoute, 0)
	if er != nil {
		return er
	}
	for _, curPilot := range pilots {
		if curPilot.ID != id {
			r.AddPilot(curPilot)
		}
	}
	return nil

}

func NewPilotSevice(fileRoute string) PilotService {
	return &pilotService{fileRoute: fileRoute}
}
