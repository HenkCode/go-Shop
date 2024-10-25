package app

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	DB 	   *gorm.DB
	Router *fiber.App
}

func (server *Server) Initialize() {
	fmt.Println("Server Go-Shop berjalan...")

	server.Router = fiber.New()
	server.initializeRoutes()
}
func (server *Server) Run(address string) {
	fmt.Printf("Berjalan di Port: %s", address)

	log.Fatal(server.Router.Listen(address))
}

func Run() {
	var server = Server{}

	server.Initialize()
	server.Run(":8080")


}
