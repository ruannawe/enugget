package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ruannawe/enugget/routes"
)

func main() {
	app := fiber.New()

	routes.setup(app)

	app.Listen(3000)
}
