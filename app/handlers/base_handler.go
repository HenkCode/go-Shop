package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/HenkCode/go-Shop/app/models"
	"github.com/HenkCode/go-Shop/database/seeders"
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AppConfig struct {
	Name string
	Env  string
	Port string
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

type Server struct {
	DB     *gorm.DB
	Router *fiber.App
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Server " + appConfig.Name + " Berjalan...")

	server.initializeDB(dbConfig)
	server.initializeRoutes()
}


func (server *Server) Run(address string) {
	fmt.Printf("Berjalan di Port: %s", address)
	
	log.Fatal(server.Router.Listen(address))
}

func (server *Server) initializeDB(dbConfig DBConfig) {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed connecting to the database server.")
	}
}

func (server *Server) dbMigrate() {
	RegisterModels := models.RegisterModels()

	for _, model := range RegisterModels {
		err := server.DB.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Succesfully Migrate...")
	}

}

func (server *Server) InitCommands(appConfig AppConfig, dbConfig DBConfig) {
	server.initializeDB(dbConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []*cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.DBSeeder(server.DB)
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}