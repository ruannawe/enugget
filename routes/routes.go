package routes

import (
	"github.com/gofiber/fiber/v2"
	usersController "github.com/ruannawe/enugget/controllers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		usersController.FindAll(c)
		return nil
	})
}
