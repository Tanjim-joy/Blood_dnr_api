package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase reads env vars and opens a MySQL connection.
// Works for Aiven, local MySQL, Railway, or any MySQL 8 server.
func ConnectDatabase() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE") // "true" for Aiven, "false" for local

	if user == "" || host == "" || name == "" {
		log.Fatal("DB_USER, DB_HOST, DB_NAME are required env vars")
	}
	if port == "" {
		port = "26274" // default Aiven port
	}
	if sslMode == "" || sslMode == "true" {
		// sslMode = "true"
		sslMode = "skip-verify"
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=%s",
		user, pass, host, port, name, sslMode,
	)

	var err error
	// Retry — Aiven sometimes needs a few seconds on cold start
	for i := 1; i <= 5; i++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Database connected")
			return
		}
		log.Printf("DB connect attempt %d/5 failed: %v", i, err)
		time.Sleep(3 * time.Second)
	}
	log.Fatal("❌ Could not connect to database: ", err)
}
