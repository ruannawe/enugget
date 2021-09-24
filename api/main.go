package main

import (
	"github.com/ruannawe/enugget/api/src/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routers.Init(app)

	app.Listen(":3000")
}
