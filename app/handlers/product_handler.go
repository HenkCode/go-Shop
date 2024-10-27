package handlers

import (
	"github.com/HenkCode/go-Shop/app/models"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) Product(c *fiber.Ctx) error {
	productModel := models.Product{}
	products, err := productModel.GetProduct(server.DB)
	if err != nil {
		return c.Status(500).SendString("Render Products Error...")
	}
	
	return c.Render("layout", fiber.Map{
		"products": products,
	})
}