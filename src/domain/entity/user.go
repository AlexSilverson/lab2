package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID      uint   `gorm:"primary_key;autoIncrement:true" `
	Login   string `gorm:"unique"`
	HashPas string
	Role    uint
}
