package app

import (
	"github.com/HenkCode/go-Shop/app/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func (server *Server) initializeRoutes() {
	engine := html.New("./templates", ".html")

	server.Router = fiber.New(fiber.Config{
		Views: engine,
	})
	
	server.Router.Static("/css", "./assets/css")
	server.Router.Static("/js", "./assets/js")
	server.Router.Static("/img", "./assets/img")
	server.Router.Static("/demo", "./assets/demo")
	server.Router.Static("/fonts", "./assets/fonts")
	server.Router.Static("/scss", "./assets/scss")

	server.Router.Get("/", handlers.Home)
}