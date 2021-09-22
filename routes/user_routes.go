package routes

import (
	"github.com/gofiber/fiber/v2"
	usersController "github.com/ruannawe/enugget/controllers"
)

func userRoutes(app *fiber.App) {
	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		usersController.FindAll(c)
		return nil
	})

	app.Get("/api/v1/user", func(c *fiber.Ctx) error {
		usersController.FindOne(c)
		return nil
	})

	app.Post("/api/v1/users", func(c *fiber.Ctx) error {
		usersController.Create(c)
		return nil
	})

	app.Patch("/api/v1/users", func(c *fiber.Ctx) error {
		usersController.Update(c)
		return nil
	})

	app.Delete("/api/v1/users", func(c *fiber.Ctx) error {
		usersController.Destroy(c)
		return nil
	})
}
