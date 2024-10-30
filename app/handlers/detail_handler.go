package handlers

import (
	"github.com/HenkCode/go-Shop/app/models"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) GetProductBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	if slug == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Slug cannot be empty")
	}

	product := models.Product{}
	foundProduct, err := product.FindBySlug(server.DB, slug)
	if err != nil {

		return c.Status(fiber.StatusNotFound).SendString("Product not found")
	}

	if foundProduct == nil {
		return c.Status(fiber.StatusNotFound).SendString("Product not found")
	}

	return c.Render("layout", fiber.Map{
		"product": foundProduct,
	})
}