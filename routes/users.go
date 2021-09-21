package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ruannawe/controllers/users"
)

func setup(app *fiber.App) {
	app.Get("/api/v1/user", users.GetUsers)
	app.Get("/api/v1/user/:id", users.GetUser)
	app.Post("/api/v1/user", users.CreateUser)
	app.Delete("/api/v1/user/:id", users.DestroyUser)
}
