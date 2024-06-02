package midlvare

import (
	"AlexSilverson/lab2/src/utility"

	"github.com/gofiber/fiber/v2"
)

func Aunthorization(c *fiber.Ctx) error {
	token := c.Get("token")
	//fmt.Println(token)
	claims, er := utility.ParseToken(token)
	if er != nil {
		return c.Status(fiber.StatusNotFound).JSON("invalid token")

	}

	if claims.Role == 0 {
		return c.Status(fiber.StatusNotFound).JSON("not enough rights")
	}

	return c.Next()
}
