package handlers

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/HenkCode/go-Shop/app/models"
	"github.com/HenkCode/go-Shop/database/seeders"
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB        *gorm.DB
	Router    *fiber.App
	AppConfig *AppConfig
}

type AppConfig struct {
	Name string
	Env  string
	Port string
	URL  string
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

type PageLink struct {
	Page 		  int32
	Url 		  string
	IsCurrentPage bool
}

type PaginationLink struct {
	CurrentPage string
	NextPage  	string
	PrevPage 	string
	TotalPage 	string
	TotalRows 	int32
	TotalPages  int32
	Links 		[]PageLink
}

type PaginationParams struct {
	Path 	  	string
	TotalRows 	int32
	PerPage	    int32
	CurrentPage int32
}


func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Server " + appConfig.Name + " Berjalan...")

	server.initializeDB(dbConfig)
	server.initializeAppConfig(appConfig)
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

func (server *Server) initializeAppConfig(appConfig AppConfig) {
	server.AppConfig = &appConfig
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

func GetPaginationLinks(appConfig *AppConfig, params PaginationParams) (PaginationLink, error) {
	var links []PageLink

	totalPages := int32(math.Ceil(float64(params.TotalRows) / float64(params.PerPage)))

	for i := 1; int32(i) <= totalPages; i++ {
		links = append(links, PageLink{
			Page: 		   int32(i),
			Url:  		   fmt.Sprintf("%s/%s?page=%s", appConfig.URL, params.Path, fmt.Sprint(i)),
			IsCurrentPage: int32(i) == params.CurrentPage,
		})
	}

	var prevPage, nextPage int32
	prevPage = 1
	nextPage = totalPages

	if params.CurrentPage > 2 {
		prevPage = params.CurrentPage - 1
	}

	if params.CurrentPage < totalPages {
		nextPage = params.CurrentPage + 1
	}

	return PaginationLink{
		CurrentPage: fmt.Sprintf("%s/%s?page=%s", appConfig.URL, params.Path, fmt.Sprint(params.CurrentPage)),
		NextPage: 	 fmt.Sprintf("%s/%s?page=%s", appConfig.URL, params.Path, fmt.Sprint(nextPage)),
		PrevPage:	 fmt.Sprintf("%s/%s?page=%s", appConfig.URL, params.Path, fmt.Sprint(prevPage)),
		TotalRows: 	 params.TotalRows,
		TotalPages:  totalPages,
		Links: 		 links,
	}, nil

}