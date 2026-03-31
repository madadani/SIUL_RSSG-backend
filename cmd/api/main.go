package main

import (
	"log"

	"siul-pbj-api/pkg/config"
	"siul-pbj-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database Connection
	config.ConnectDB()

	// Initialize Gin router
	r := gin.Default()

	// CORS Setup
	configCORS := cors.DefaultConfig()
	configCORS.AllowAllOrigins = true
	configCORS.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(configCORS))

	// Register Routes (inject DB)
	routes.SetupRoutes(r, config.DB)

	// Start Server
	log.Println("Server running on port 8080")
	r.Run(":8080")
}
