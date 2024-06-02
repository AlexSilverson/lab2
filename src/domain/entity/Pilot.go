package entity

import "gorm.io/gorm"

type Pilot struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primary_key;autoIncrement:true"`
	FirstName  string `json:"firstname"`
	SecondName string `json:"secondname"`
	LastName   string `json:"lastname"`
	Age        uint   `json:"age"`
	Experience uint   `json:"experience"`
}
