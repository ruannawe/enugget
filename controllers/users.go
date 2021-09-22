package controllers

import (
	"github.com/gofiber/fiber/v2"
	usersService "github.com/ruannawe/enugget/services/users"
)

func FindAll(c *fiber.Ctx) {
	c.SendString(usersService.FindAll())
}

func FindOne(c *fiber.Ctx) {
	usersService.FindOne()
}

func Create(c *fiber.Ctx) {
	usersService.Create()
}

func Update(c *fiber.Ctx) {
	usersService.Update()
}

func Destroy(c *fiber.Ctx) {
	usersService.Destroy()
}
