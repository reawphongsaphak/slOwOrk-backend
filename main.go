package main

import (
	"log"
	
	"github.com/reawphongsaphak/slOwOrk-Backend/config"
	"github.com/reawphongsaphak/slOwOrk-Backend/database"
	"github.com/reawphongsaphak/slOwOrk-Backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	cfg := config.LoadConfig()

	database.ConnectDB()

	// Initialize Gin
	r := gin.Default()
	r.Use(cors.Default())
	
	routes.SetupRoutes(r)

	r.Run(cfg.Port)
}