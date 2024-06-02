package entity

import "gorm.io/gorm"

type City struct {
	gorm.Model
	ID      uint   `json:"id" gorm:"primary_key;autoIncrement:true" validate:"required"`
	Name    string `json:"name" validate:"required" `
	Country string `json:"country" validate:"required"`
	History uint   `json:"history" validate:"required" `
	Resort  uint   `json:"resort" validate:"required" `
}
