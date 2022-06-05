package main

// go:generate sqlboiler --wipe psql

import (
	"go-prima/handler"
	"go-prima/repository"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/joho/godotenv"

	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get Config
	config, err := getConfig()
	if err != nil {
		panic(err)
	}

	dbURL := getDBSourceURL(&config.DBConfig)

	conn, err := sql.Open(os.Getenv("DB_DRIVER"), dbURL)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := repository.NewStore(conn)

	application := handler.NewServer(app, store)

	fmt.Println(config.AppConfig.AppName + " Open Port at " + config.AppConfig.AppPort)
	application.Listen(":" + config.AppConfig.AppPort)
}

func getDBSourceURL(config *DBConfig) string {

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
}
