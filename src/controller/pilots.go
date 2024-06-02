package controller

import (
	"AlexSilverson/lab2/src/domain/entity"
	"AlexSilverson/lab2/src/domain/services"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetPilotById Getting Pilot by ID
//
//	@Summary		Getting Pilots by Id
//	@Description	Getting Pilots by Id in detail
//	@Tags			Pilots
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"id of Pilot"
//	@Param			token			header		string	true	"auth token"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/pilot/{id} [get]
func GetPilotById(app *fiber.App, pilotService services.PilotService) fiber.Router {
	return app.Get("/pilot/:id", func(c *fiber.Ctx) error {
		pilotId := c.Params("id")
		id, err := strconv.ParseInt(pilotId, 10, 64)
		fmt.Println("here ", c.Get("token"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Id format is not valid")
		}

		pilot, er := pilotService.GetPilotById(uint(id))
		if er != nil {
			return c.Status(fiber.StatusNotFound).JSON(string(er.Error()))
		}

		return c.Status(fiber.StatusOK).JSON(pilot)
	})
}

// PutPilotById Getting Pilot by json
//
//	@Summary		Getting Pilot by Id
//	@Description	Getting Pilot by Id in detail
//	@Tags			Pilots
//	@Accept			json
//	@Produce		json
//	@Param			request			body		entity.Pilot	true	"Request of Creating Pilot Object"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/pilot [post]
func AddPilot(app *fiber.App, pilotService services.PilotService) fiber.Router {
	return app.Post("/pilot", func(c *fiber.Ctx) error {
		var pilot entity.Pilot
		err := c.BodyParser(&pilot)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Pilot format is not valid")
		}
		err = pilotService.AddPilot(pilot)
		if err == nil {
			return c.Status(fiber.StatusOK).JSON("added")
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

	})
}

// GetPilotById Updating Pilot
//
//	@Summary		Updating Pilot
//	@Description	Updating Pilot in detail
//	@Tags			Pilots
//	@Accept			json
//	@Produce		json
//	@Param			request			body		entity.Pilot	true	"Request of Updating Pilot Object"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/pilot [put]
func UpdatePilot(app *fiber.App, pilotService services.PilotService) fiber.Router {
	return app.Put("/pilot", func(c *fiber.Ctx) error {
		var pilot entity.Pilot

		err := c.BodyParser(&pilot)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Pilot format is not valid")
		}

		er := pilotService.UpdatePilot(pilot)
		if er != nil {
			return c.Status(fiber.StatusBadRequest).JSON(er.Error())
		}
		return c.Status(fiber.StatusOK).JSON("updated")
	})
}

// DeletePilot Deliting Pilot by ID
//
//	@Summary		Deletting Pilot by Id
//	@Description	Deletting Pilot by Id in detail
//	@Tags			Pilots
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"id of Pilot"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/pilot/{id} [delete]
func DeletePilot(app *fiber.App, pilotService services.PilotService) fiber.Router {

	return app.Delete("/pilot/:id", func(c *fiber.Ctx) error {
		pilotId := c.Params("id")
		id, err := strconv.ParseInt(pilotId, 10, 64)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		er := pilotService.DeletePilot(uint(id))
		if er != nil {
			return c.Status(fiber.StatusNotFound).JSON(string(er.Error()))
		}

		return c.Status(fiber.StatusOK).JSON("deleted")
	})
}
