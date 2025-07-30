package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/khanjaved9700/todo_app/config"
	"github.com/khanjaved9700/todo_app/routes"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := config.InitDB()

	r := gin.Default()

	routes.RegisterRoutes(r, db)
	r.Run(":8080")
}
