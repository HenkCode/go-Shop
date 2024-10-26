package app

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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
	fmt.Println("DSN:", dsn)
	server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connecting to the database server.")
	}
	
	for _, model := range RegisterModels() {
		err = server.DB.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database migrated successfully.")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on Loading .env file.")
	}

	appConfig.Name = getEnv("APP_NAME", "go-Shop")
	appConfig.Env  = getEnv("APP_ENV", "dev")
	appConfig.Port = getEnv("APP_PORT", "9000")

	dbConfig.Host 	  = getEnv("DB_HOST", "localhost")
	dbConfig.Name     = getEnv("DB_NAME", "go-shop")
	dbConfig.User     = getEnv("DB_USER", "root")
	dbConfig.Password = getEnv("DB_PASSWORD", "")
	dbConfig.Port     = getEnv("DB_PORT", "5432")

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.Port)
}
