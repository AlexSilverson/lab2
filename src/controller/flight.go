package controller

import (
	"AlexSilverson/lab2/src/domain/entity"
	"AlexSilverson/lab2/src/domain/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetFlightById Getting Flight by ID
//
//	@Summary		Getting Flight by Id
//	@Description	Getting Flight by Id in detail
//	@Tags			Flights
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"id of Flight"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/flight/{id} [get]
func GetFlightById(app *fiber.App, flightService services.FlightService) fiber.Router {
	return app.Get("/flight/:id", func(c *fiber.Ctx) error {
		flightId := c.Params("id")
		id, err := strconv.ParseInt(flightId, 10, 64)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Id format is not valid")
		}

		flight, er := flightService.GetFlightById(uint(id))
		if er != nil {
			return c.Status(fiber.StatusNotFound).JSON(er)
		}

		return c.Status(fiber.StatusOK).JSON(flight)
	})
}

// PutFlightById Getting Flight by json
//
//	@Summary		Getting Flight by Id
//	@Description	Getting Flight by Id in detail
//	@Tags			Flights
//	@Accept			json
//	@Produce		json
//	@Param			request			body		entity.Flight	true	"Request of Creating Flight Object"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/flight [post]
func AddFlight(app *fiber.App, flightService services.FlightService) fiber.Router {
	return app.Post("/flight", func(c *fiber.Ctx) error {
		var flight entity.Flight
		err := c.BodyParser(&flight)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Flight format is not valid")
		}
		_, er := flightService.GetFlightById(flight.ID)
		if er != nil && er.Error() == "flight not found" {
			err = flightService.AddFlight(flight)
			if err == nil {
				return c.Status(fiber.StatusOK).JSON("added")
			} else {
				return c.Status(fiber.StatusBadRequest).JSON(err)
			}
		}

		return c.Status(fiber.StatusBadRequest).JSON("That id is already used")

	})
}

// GetFlightById Updating Flight
//
//	@Summary		Updating Flight
//	@Description	Updating Flight in detail
//	@Tags			Flights
//	@Accept			json
//	@Produce		json
//	@Param			request			body		entity.Flight	true	"Request of Updating Flight Object"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/flight [put]
func UpdateFlight(app *fiber.App, flightService services.FlightService) fiber.Router {
	return app.Put("/flight", func(c *fiber.Ctx) error {
		var flight entity.Flight

		err := c.BodyParser(&flight)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Flight format is not valid")
		}
		_, er := flightService.GetFlightById(flight.ID)

		if er != nil {
			return c.Status(fiber.StatusBadRequest).JSON(er.Error())
		}
		er = flightService.UpdateFlight(flight)
		if er != nil {
			return c.Status(fiber.StatusBadRequest).JSON(er.Error())
		}
		return c.Status(fiber.StatusOK).JSON("updated")
	})
}

// DeleteFlight Deliting Flightt by ID
//
//	@Summary		Deletting Flight by Id
//	@Description	Deletting Flight by Id in detail
//	@Tags			Flights
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"id of Flight"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/flight/{id} [delete]
func DeleteFlight(app *fiber.App, flightService services.FlightService) fiber.Router {

	return app.Delete("/flight/:id", func(c *fiber.Ctx) error {
		flightId := c.Params("id")
		id, err := strconv.ParseInt(flightId, 10, 64)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		er := flightService.DeleteFlight(uint(id))
		if er != nil {
			return c.Status(fiber.StatusNotFound).JSON("Flight not found")
		}

		return c.Status(fiber.StatusOK).JSON("deleted")
	})
}
