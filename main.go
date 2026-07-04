package main

import (
	"log"
	"os"

	"blood-donor-api/config"
	"blood-donor-api/handlers"
	"blood-donor-api/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to system env")
	}

	// Connect to MySQL (Aiven / local / Railway — same code)
	config.ConnectDatabase()

	// Auto migrate tables
	if err := config.DB.AutoMigrate(&models.Donor{}, &models.BloodRequest{}); err != nil {
		log.Fatal("Migration failed: ", err)
	}

	// Setup router
	r := gin.Default()

	// CORS — open in dev, restrict in prod
	corsOrigins := os.Getenv("CORS_ORIGINS")
	if corsOrigins == "" {
		corsOrigins = "*"
	}
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{corsOrigins},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// Health check
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "blood-donor-api"})
	})

	api := r.Group("/api")
	{
		// Donors
		api.GET("/donors", handlers.GetDonors)
		api.GET("/donors/:id", handlers.GetDonor)
		api.POST("/donors", handlers.CreateDonor)
		api.PUT("/donors/:id", handlers.UpdateDonor)
		api.DELETE("/donors/:id", handlers.DeleteDonor)

		// Blood requests
		api.GET("/requests", handlers.GetRequests)
		api.POST("/requests", handlers.CreateRequest)
		api.PUT("/requests/:id/status", handlers.UpdateRequestStatus)
		api.DELETE("/requests/:id", handlers.DeleteRequest)

		// Dashboard stats
		api.GET("/stats", handlers.GetStats)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🩸 Blood donor API running on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server failed: ", err)
	}
}
