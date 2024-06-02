package controller

import (
	"AlexSilverson/lab2/src/domain/entity"
	"AlexSilverson/lab2/src/domain/entity/dtos"
	"AlexSilverson/lab2/src/domain/services"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// AddUser Adding User by json
//
//	@Summary		Add User by json
//	@Description	Add User by jso in detail
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request			body		dtos.UserDto	true	"Request of Creating User Object"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/user [post]
func AddUser(app *fiber.App, userService services.UserService) fiber.Router {
	return app.Post("/user", func(c *fiber.Ctx) error {
		var UserDto dtos.UserDto
		var User entity.User
		err := c.BodyParser(&UserDto)
		User = UserDto.MapNewUserDtoToUser()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("User format is not valid")
		}
		err = userService.AddUser(User)
		if err == nil {
			return c.Status(fiber.StatusOK).JSON("added")
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}
	})
}

// Login User Adding User by json
//
//		@Summary		Login User by json
//		@Description	Login User by jso in detail
//		@Tags			Users
//		@Accept			json
//		@Produce		json
//		@Param			login			header		string	true	"Login"
//	 	@Param			password		header		string	true	"Password"
//		@Failure		400				{string}	string
//		@Failure		404				{string}	string
//		@Success		200				{string}	string
//		@Router			/login [get]
func LoginUser(app *fiber.App, userService services.UserService) fiber.Router {
	return app.Get("/login", func(c *fiber.Ctx) error {
		var userDto dtos.UserDto
		//fmt.Println("here ", c.Params("login"))
		userDto.Login = c.Get("login")
		userDto.Password = c.Get("password")
		fmt.Println(userDto)
		var user entity.User = userDto.MapNewUserDtoToUser()
		fmt.Println("here ", user)
		token, err := userService.Login(user)
		if err == nil {
			return c.Status(fiber.StatusOK).JSON(token)
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
	})
}

// GetUserByLogin Getting User by Login
//
//	@Summary		Getting User by Login
//	@Description	Getting User by Login in detail
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"id of User"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/user/{id} [get]
func GetUserById(app *fiber.App, userService services.UserService) fiber.Router {
	return app.Get("/user/:id", func(c *fiber.Ctx) error {
		userId := c.Params("id")
		id, err := strconv.ParseInt(userId, 10, 64)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Id format is not valid")
		}

		city, er := userService.GetUserById(uint(id))
		if er != nil {
			return c.Status(fiber.StatusNotFound).JSON(string(er.Error()))
		}

		return c.Status(fiber.StatusOK).JSON(city)
	})
}

// UpdatingUserRole Updating User Role
//
//	@Summary		Updating User Role
//	@Description	Updating User Role in detail
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id			header		string 	true	"Request of Updating Flight Object"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/user [put]
func UpdateUserRole(app *fiber.App, userService services.UserService) fiber.Router {
	return app.Put("/user", func(c *fiber.Ctx) error {

		id, er := strconv.ParseInt(c.Get("id"), 10, 64)

		if er != nil {
			return c.Status(fiber.StatusBadRequest).JSON("Id format is not valid")
		}

		er = userService.UpdRole(uint(id))
		if er != nil {
			return c.Status(fiber.StatusBadRequest).JSON(er.Error())
		}
		return c.Status(fiber.StatusOK).JSON("updated")
	})
}
