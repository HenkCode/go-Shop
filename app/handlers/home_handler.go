package handlers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.SendString("i'm home page")
}