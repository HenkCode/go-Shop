package app

import (
	"flag"
	"log"
	"os"

	"github.com/HenkCode/go-Shop/app/handlers"
	"github.com/joho/godotenv"
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func Run() {
	var server = handlers.Server{}
	var appConfig = handlers.AppConfig{}
	var dbConfig = handlers.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on Loading .env file.")
	}

	appConfig.Name = getEnv("APP_NAME", "go-Shop")
	appConfig.Env = getEnv("APP_ENV", "dev")
	appConfig.Port = getEnv("APP_PORT", "9000")

	dbConfig.Host = getEnv("DB_HOST", "localhost")
	dbConfig.Name = getEnv("DB_NAME", "go-shop")
	dbConfig.User = getEnv("DB_USER", "root")
	dbConfig.Password = getEnv("DB_PASSWORD", "")
	dbConfig.Port = getEnv("DB_PORT", "5432")

	flag.Parse()
	arg := flag.Arg(0)

	if arg != "" {
		server.InitCommands(appConfig, dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.Port)
	}
}
