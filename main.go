package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/your_username/your_project/handlers"
	"github.com/your_username/your_project/middleware"
)

func main() {
	r := gin.Default()

	// Добавление CORS middleware
	r.Use(middleware.CORSMiddleware())

	// Настройки подключения к базе данных
	databaseURLChinook := os.Getenv("DATABASE_URL_CHINOOK")
	if databaseURLChinook == "" {
		databaseURLChinook = "postgresql://postgres:postgres@localhost/chinook?sslmode=disable"
	}

	databaseURLSakila := os.Getenv("DATABASE_URL_SAKILA")
	if databaseURLSakila == "" {
		databaseURLSakila = "postgresql://postgres:postgres@localhost/sakila?sslmode=disable"
	}

	databaseURLNorthwind := os.Getenv("DATABASE_URL_NORTHWIND")
	if databaseURLNorthwind == "" {
		databaseURLNorthwind = "postgresql://postgres:postgres@localhost/northwind?sslmode=disable"
	}

	r.POST("/chinook/query", handlers.HandleQuery(databaseURLChinook))
	r.POST("/sakila/query", handlers.HandleQuery(databaseURLSakila))
	r.POST("/northwind/query", handlers.HandleQuery(databaseURLNorthwind))

	// Добавьте другие базы данных по аналогии

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
