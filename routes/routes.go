package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ruannawe/enugget/routes/users"
)

func setup(app *fiber.App) {
	user.routes(app)
}
