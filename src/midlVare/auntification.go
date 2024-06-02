package midlvare

import (
	"AlexSilverson/lab2/src/utility"

	"github.com/gofiber/fiber/v2"
)

func Auntification(c *fiber.Ctx) error {

	token := c.Get("token")
	//fmt.Println(token)
	_, er := utility.ParseToken(token)
	if er != nil {
		return c.Status(fiber.StatusNotFound).JSON("invalid token")

	}
	return c.Next()
}
