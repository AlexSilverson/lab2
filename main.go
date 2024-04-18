package main

import (
	_ "AlexSilverson/lab2/docs"
	"AlexSilverson/lab2/src/City/controller"
	"AlexSilverson/lab2/src/City/domain/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title			2lab API
// @version		1.0
// @description	This is a sample swagger for second RCSP lab
// @termsOfService	http://swagger.io/terms/
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Delete("/swagger/del/*", swagger.HandlerDefault)
	file := "C:\\Lab2\\Storadges\\CityStorage.txt"

	cityServise := services.NewCitySevice(file)
	//fmt.Print(cityRepository.GetCityById(2))
	controller.GetCityById(app, cityServise)
	controller.AddCity(app, cityServise)
	controller.UpdateCity(app, cityServise)
	controller.DeleteCity(app, cityServise)
	//fmt.Print(cityServise.GetCityById(2))
	app.Listen(":3001")
}
