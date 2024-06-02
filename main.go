package main

import (
	"AlexSilverson/lab2/Storadges/database"
	_ "AlexSilverson/lab2/docs"
	"AlexSilverson/lab2/src/controller"
	midlvare "AlexSilverson/lab2/src/midlVare"

	"AlexSilverson/lab2/src/domain/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title			2lab API
// @version		1.0
// @description	This is a sample swagger for second RCSP lab
// @termsOfService	http://swagger.io/terms/
func main() {
	db := database.Init()

	//fmt.Println(ans)
	//token, _ := jwt.ParseSigned(ans)

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	auntGroup := app.Group("/aunt")
	authGroup := app.Group("/auth")
	//publicGroup := app.Group("/public")

	auntGroup.Use(midlvare.Auntification)

	authGroup.Use(midlvare.Auntification)
	authGroup.Use(midlvare.Aunthorization)

	cityServise := services.NewCitySevice(db)
	pilotService := services.NewPilotSevice(db)
	//flightFile := "C:\\Lab2\\Storadges\\FlightStorage.txt"
	flightService := services.NewFlightSevice(db)
	userService := services.NewUserSevice(db)
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
	controller.GetShortFlightById(app, flightService)

	controller.AddUser(app, userService)
	controller.LoginUser(app, userService)
	controller.GetUserById(app, userService)
	controller.UpdateUserRole(app, userService)
	app.Listen(":3001")

}

/*
{
  "from": 1,
  "id": 0,
  "pilot": 1,
  "starttime": "never",
  "to": 2,
  "treveltime": 0
}
*/
