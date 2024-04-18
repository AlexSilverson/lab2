package controller

import (
	"AlexSilverson/lab2/src/domain/entity"
	"AlexSilverson/lab2/src/domain/services"
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
		cityId := c.Params("id")
		id, err := strconv.ParseInt(cityId, 10, 64)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Id format is not valid")
		}

		city, er := cityService.GetCityById(uint(id))
		if er != nil {
			return c.Status(fiber.StatusNotFound).JSON(er)
		}

		return c.Status(fiber.StatusOK).JSON(city)
	})
}

// PutCityById Getting City by json
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
		_, er := cityService.GetCityById(city.ID)
		if er != nil && er.Error() == "city not found" {
			err = cityService.AddCity(city)
			if err == nil {
				return c.Status(fiber.StatusOK).JSON("added")
			} else {
				return c.Status(fiber.StatusBadRequest).JSON(err)
			}
		}

		return c.Status(fiber.StatusBadRequest).JSON("That id is already used")

	})
}

// GetCityById Updating City
//
//	@Summary		Updating City
//	@Description	Updating City in detail
//	@Tags			Citis
//	@Accept			json
//	@Produce		json
//	@Param			request			body		entity.City	true	"Request of Updating City Object"
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
		_, er := cityService.GetCityById(city.ID)

		if er != nil {
			return c.Status(fiber.StatusBadRequest).JSON(er.Error())
		}
		er = cityService.UpdateCity(city)
		if er != nil {
			return c.Status(fiber.StatusBadRequest).JSON(er.Error())
		}
		return c.Status(fiber.StatusOK).JSON("updated")
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
		cityId := c.Params("id")
		id, err := strconv.ParseInt(cityId, 10, 64)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		er := cityService.DeleteCity(uint(id))
		if er != nil {
			return c.Status(fiber.StatusNotFound).JSON("City not found")
		}

		return c.Status(fiber.StatusOK).JSON("deleted")
	})
}
