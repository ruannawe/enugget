package controllers

import (
	"github.com/ruannawe/enugget/api/src/services"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) {
	c.Set("Content-type", "application/json; charset=utf-8")
	c.Send(services.GetUsers())
}

func GetUser(c *fiber.Ctx) {
	c.Set("Content-type", "application/json; charset=utf-8")
	services.GetUser()
}

func CreateUser(c *fiber.Ctx) {
	c.Set("Content-type", "application/json; charset=utf-8")
	services.CreateUser()
}

func UpdateUser(c *fiber.Ctx) {
	c.Set("Content-type", "application/json; charset=utf-8")
	services.UpdateUser()
}

func DestroyUser(c *fiber.Ctx) {
	c.Set("Content-type", "application/json; charset=utf-8")
	services.DestroyUser()
}
