package main

import (
	"agenda-personal/database"
	"agenda-personal/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConectandoDB()
	router.IniciandoRotas(app)

	app.Listen(":3000")
}