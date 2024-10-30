package handlers

import "github.com/gofiber/fiber/v2"

func (Server *Server) Home(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title": "Fast Shipping",
	})
}