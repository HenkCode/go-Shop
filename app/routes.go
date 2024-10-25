package app

import "github.com/HenkCode/go-Shop/app/handlers"

func (server *Server) initializeRoutes() {
	server.Router.Get("/", handlers.Home)
}