package routers

import (
	"github.com/ruannawe/enugget/api/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func userRoutes(app *fiber.App) {
	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		controller.GetUsers(c)
		return nil
	})

	app.Get("/api/v1/users/:id", func(c *fiber.Ctx) error {
		controller.GetUser(c)
		return nil
	})

	app.Post("/api/v1/users", func(c *fiber.Ctx) error {
		controller.CreateUser(c)
		return nil
	})

	app.Patch("/api/v1/users", func(c *fiber.Ctx) error {
		controller.UpdateUser(c)
		return nil
	})

	app.Delete("/api/v1/users", func(c *fiber.Ctx) error {
		controller.DestroyUser(c)
		return nil
	})
}
