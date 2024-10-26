package app

import (
	"github.com/HenkCode/go-Shop/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func (server *Server) initializeRoutes() {
	server.Router = fiber.New()
	
	server.Router.Get("/", handlers.Home)
}