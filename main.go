package main

import (
	_ "AlexSilverson/lab2/docs"
	"AlexSilverson/lab2/src/controller"

	"AlexSilverson/lab2/src/domain/services"

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
	cityFile := "C:\\Lab2\\Storadges\\CityStorage.txt"
	cityServise := services.NewCitySevice(cityFile)
	pilotFile := "C:\\Lab2\\Storadges\\PilotStorage.txt"
	pilotService := services.NewPilotSevice(pilotFile)
	flightFile := "C:\\Lab2\\Storadges\\FlightStorage.txt"
	flightService := services.NewFlightSevice(flightFile)
	//fmt.Print(cityRepository.GetCityById(2))
	controller.GetCityById(app, cityServise)
	controller.AddCity(app, cityServise)
	controller.UpdateCity(app, cityServise)
	controller.DeleteCity(app, cityServise)
	//fmt.Print(cityServise.GetCityById(2))
	controller.GetPilotById(app, pilotService)
	controller.UpdatePilot(app, pilotService)
	controller.AddPilot(app, pilotService)
	controller.DeletePilot(app, pilotService)

	controller.GetFlightById(app, flightService)
	controller.AddFlight(app, flightService, cityServise, pilotService)
	controller.UpdateFlight(app, flightService, cityServise, pilotService)
	controller.DeleteFlight(app, flightService)
	app.Listen(":3001")
}
