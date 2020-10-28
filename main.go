package main

import (
	"log"

	"github.com/hunzo/go-fiber-gorm-102/models"

	"github.com/gofiber/fiber/v2"
	"github.com/hunzo/go-fiber-gorm-102/routers"
	"github.com/hunzo/go-fiber-gorm-102/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initdb() {
	var err error
	services.DBConnect, err = gorm.Open(sqlite.Open("./database/sessionToken.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	log.Println("Connect Database")

	services.DBConnect.AutoMigrate(&models.TOKEN{})
	log.Println("Database Migrated !!")

}

func main() {
	app := fiber.New()
	initdb()

	routers.SetupRouters(app)

	app.Listen(":8080")
}
