package main

import (
	"log"

	"github.com/ahmed-afzal1/go-crud/config"
	"github.com/ahmed-afzal1/go-crud/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()

	config.ConnectDatabase()
	routes.RegisterRoutes(r)

	r.Run()
}
