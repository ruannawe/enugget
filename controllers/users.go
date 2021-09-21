package controllers

import (
	"github.com/gofiber/fiber/v2"
	usersService "github.com/ruannawe/services/users"
)

func FindAll(c *fiber.Ctx) {
	return c.SendString(usersService.FindAll())
}

func FindOne(c *fiber.Ctx) {
	return c.SendString(usersService.FindOne())
}

func Create(c *fiber.Ctx) {
	return c.SendString(usersService.Create())
}

func Update(c *fiber.Ctx) {
	return c.SendString(usersService.Update())
}

func Destroy(c *fiber.Ctx) {
	return c.SendString(usersService.Destroy())
}
