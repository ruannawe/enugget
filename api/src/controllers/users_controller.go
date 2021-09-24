package controllers

import (
	"github.com/ruannawe/enugget/api/src/services"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) {
	c.JSON(services.GetUsers())
}

func GetUser(c *fiber.Ctx) {
	services.GetUser()
}

func CreateUser(c *fiber.Ctx) {
	services.CreateUser()
}

func UpdateUser(c *fiber.Ctx) {
	services.UpdateUser()
}

func DestroyUser(c *fiber.Ctx) {
	services.DestroyUser()
}
