package middlewares

import (
	"log"

	"github.com/hunzo/go-fiber-gorm-102/services"

	"github.com/gofiber/fiber/v2"
)

func DBCheck(c *fiber.Ctx) error {
	log.Println("test middleware")
	ret, err := services.CheckData()
	if err != nil {
		return c.JSON(fiber.Map{
			"error": "record not found",
		})
	}
	log.Println(ret)
	return c.Next()
}
