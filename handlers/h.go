package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hunzo/go-fiber-gorm-102/services"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Home")
}

func Create(c *fiber.Ctx) error {
	services.CreateData()
	return c.SendString("Create Records")
}

func Delete(c *fiber.Ctx) error {
	services.DeleteData()
	return c.SendString("Delete Records")
}
