package database

import (
	"AlexSilverson/lab2/src/domain/entity"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=Alex0910? dbname=3labRksp port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("ok")
	err = db.AutoMigrate(&entity.City{}, &entity.Flight{}, &entity.Pilot{}, &entity.User{})
	if err != nil {
		panic(err)
	}
	fmt.Println("ok")
	return db
}
