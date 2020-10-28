package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hunzo/go-fiber-gorm-102/handlers"
	"github.com/hunzo/go-fiber-gorm-102/middlewares"
)

func SetupRouters(r *fiber.App) {

	// r.Use(middlewares.DBCheck)
	r.Get("/create", handlers.Create)
	r.Get("/delete", handlers.Delete)

	auth := r.Group("/auth")
	auth.Use(middlewares.DBCheck)

	auth.Get("/", handlers.Home)

}
