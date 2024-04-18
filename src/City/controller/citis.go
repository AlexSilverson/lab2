package controller

import (
	"AlexSilverson/lab2/src/City/domain/entity"
	"AlexSilverson/lab2/src/City/domain/services"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetCityById Getting City by ID
//
//	@Summary		Getting City by Id
//	@Description	Getting City by Id in detail
//	@Tags			Citis
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"id of City"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/city/{id} [get]
func GetCityById(app *fiber.App, cityService services.CityService) fiber.Router {

	return app.Get("/city/:id", func(c *fiber.Ctx) error {
		fmt.Print("here")
		cityId := c.Params("id")
		id, err := strconv.ParseInt(cityId, 10, 64)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Id format is not valid")
		}

		city := cityService.GetCityById(uint(id))
		if city == nil {
			return c.Status(fiber.StatusNotFound).JSON("City not found")
		}

		return c.Status(fiber.StatusOK).JSON(city)
	})
}

// GetCityById Getting City by ID
//
//	@Summary		Getting City by Id
//	@Description	Getting City by Id in detail
//	@Tags			Citis
//	@Accept			json
//	@Produce		json
//	@Param			request			body		entity.City	true	"Request of Creating City Object"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/city [post]
func AddCity(app *fiber.App, cityService services.CityService) fiber.Router {
	return app.Post("/city", func(c *fiber.Ctx) error {
		var city entity.City

		err := c.BodyParser(&city)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("City format is not valid")
		}

		createdCity := cityService.AddCity(city)

		return c.Status(fiber.StatusOK).JSON(createdCity)
	})
}

// GetCityById Getting City by ID
//
//	@Summary		Getting City by Id
//	@Description	Getting City by Id in detail
//	@Tags			Citis
//	@Accept			json
//	@Produce		json
//	@Param			request			body		entity.City	true	"Request of Creating City Object"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/city [put]
func UpdateCity(app *fiber.App, cityService services.CityService) fiber.Router {
	return app.Put("/city", func(c *fiber.Ctx) error {
		var city entity.City

		err := c.BodyParser(&city)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("City format is not valid")
		}

		createdCity := cityService.UpdateCity(city)

		return c.Status(fiber.StatusOK).JSON(createdCity)
	})
}

/*
{
  "country": "France",
  "history": 9,
  "id": 4,
  "name": "Paris",
  "resort": 100
}
*/

// GetCityById Getting City by ID
//
//	@Summary		Deletting City by Id
//	@Description	Deletting City by Id in detail
//	@Tags			Citis
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"id of City"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/city/{id} [delete]
func DeleteCity(app *fiber.App, cityService services.CityService) fiber.Router {

	return app.Delete("/city/:id", func(c *fiber.Ctx) error {
		fmt.Print("here")
		cityId := c.Params("id")
		id, err := strconv.ParseInt(cityId, 10, 64)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Id format is not valid")
		}
		fmt.Println(id)

		city := cityService.DeleteCity(uint(id))
		fmt.Println("delete: ", id, " ", city)
		if city == nil {
			return c.Status(fiber.StatusNotFound).JSON("City not found")
		}

		return c.Status(fiber.StatusOK).JSON(city)
	})
}
