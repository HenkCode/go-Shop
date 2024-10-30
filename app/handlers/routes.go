package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func (server *Server) initializeRoutes() {
	engine := html.New("./templates", ".html")

	server.Router = fiber.New(fiber.Config{
		Views: engine,
		ViewsLayout: "layout",
		PassLocalsToViews: true,
	})
	
	staticDirs := map[string]string{
        "/css":   "./assets/css",
        "/js":    "./assets/js",
        "/img":   "./assets/img",
        "/demo":  "./assets/demo",
        "/fonts": "./assets/fonts",
        "/scss":  "./assets/scss",
    }

    for route, path := range staticDirs {
        server.Router.Static(route, path)
    }

	server.Router.Get("/", server.Home)
	server.Router.Get("/products", server.Products)
	server.Router.Get("/products/:slug", server.GetProductBySlug)
}