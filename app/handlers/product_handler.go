package handlers

import (
	"fmt"
	"strconv"

	"github.com/HenkCode/go-Shop/app/models"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) Products(c *fiber.Ctx) error {
	Query := c.Query("page", "1")
	page, err := strconv.Atoi(Query)

	fmt.Println("Received page query:", Query)

	if err != nil || page <= 0 {
		page = 1
	}

	perPage := 9

	productModel := models.Product{}
	products, totalRows, err := productModel.GetProduct(server.DB, perPage, page)
	if err != nil {
		fmt.Println("Error fetching products:", err)
		return c.Status(500).SendString("Render Products Error...")
	}
	fmt.Println("totalrows: ", totalRows)

	pagination, err := GetPaginationLinks(server.AppConfig, PaginationParams{
		Path: 	     "products",
		TotalRows:   int32(totalRows),
		PerPage:     int32(perPage),
		CurrentPage: int32(page),
	})
	if err != nil {
		fmt.Println("Error getting pagination links:", err) // Log kesalahan pagination
        return c.Status(500).SendString("Pagination Error...")
	}
	fmt.Println("Pagination Links: ", pagination)
	
	return c.Render("layout", fiber.Map{
		"products": products,
		"pagination": pagination,
	})
}
