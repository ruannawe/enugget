package main

import (
	usersController "github.com/ruannawe/enugget/controllers"
	// teste "github.com/ruannawe/enugget/routes"
	"github.com/gofiber/fiber/v2"
	"fmt"
)

func main() {
	app := fiber.New()

	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		usersController.FindAll(c)
		return nil
	})

	app.Get("/api/v1/user", func(c *fiber.Ctx) error {
		usersController.FindOne(c)
		return nil
	})

	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		usersController.Update(c)
		return nil
	})

	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		usersController.Destroy(c)
		return nil
	})

	app.Listen(":3000")
}
