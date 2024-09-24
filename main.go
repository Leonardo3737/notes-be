package main

import (
	"log"
	"os"

	"github.com/Leonardo3737/notes-be/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
