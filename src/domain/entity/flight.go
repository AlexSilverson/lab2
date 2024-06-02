package entity

import (
	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	ID         uint `gorm:"primary_key;autoIncrement:true" `
	FromId     uint
	From       City `gorm:"foreignKey:FromId"`
	ToId       uint
	To         City `gorm:"foreignKey:ToId"`
	TrevelTime uint ``
	PilotId    uint
	Pilot      Pilot  `gorm:"foreignKey:PilotId"`
	StartTime  string ``
}
