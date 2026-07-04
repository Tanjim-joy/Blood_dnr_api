package config

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getenv(keys ...string) string {
	for _, key := range keys {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}
	return ""
}

func parseDatabaseURL(raw string) (user, pass, host, port, name, sslMode string, err error) {
	u, err := url.Parse(raw)
	if err != nil {
		return
	}

	if u.User != nil {
		user = u.User.Username()
		pass, _ = u.User.Password()
	}

	host = u.Hostname()
	port = u.Port()
	if port == "" {
		port = "21274" // default MySQL port
	}

	name = strings.TrimPrefix(u.Path, "/")
	if name == "" {
		err = fmt.Errorf("database name missing in URL")
		return
	}

	query := u.Query()
	sslMode = query.Get("tls")
	if sslMode == "" {
		sslMode = query.Get("sslmode")
	}

	return
}

// ConnectDatabase reads env vars and opens a MySQL connection.
// Works for Aiven, local MySQL, Railway, or any MySQL 8 server.
func ConnectDatabase() {
	user := getenv("DB_USER", "MYSQL_USER", "DATABASE_USER", "MYSQL_USERNAME")
	pass := getenv("DB_PASSWORD", "MYSQL_PASSWORD", "DATABASE_PASSWORD")
	host := getenv("DB_HOST", "MYSQL_HOST", "DATABASE_HOST")
	port := getenv("DB_PORT", "MYSQL_PORT", "DATABASE_PORT")
	name := getenv("DB_NAME", "MYSQL_DATABASE", "DATABASE_NAME")
	sslMode := getenv("DB_SSL_MODE", "MYSQL_SSL_MODE")

	if user == "" || host == "" || name == "" {
		if dbURL := getenv("DATABASE_URL", "MYSQL_URL", "MYSQL_DSN"); dbURL != "" {
			parsedUser, parsedPass, parsedHost, parsedPort, parsedName, parsedSSL, err := parseDatabaseURL(dbURL)
			if err == nil {
				if user == "" {
					user = parsedUser
				}
				if pass == "" {
					pass = parsedPass
				}
				if host == "" {
					host = parsedHost
				}
				if port == "" {
					port = parsedPort
				}
				if name == "" {
					name = parsedName
				}
				if sslMode == "" {
					sslMode = parsedSSL
				}
			}
		}
	}

	if user == "" || host == "" || name == "" {
		log.Fatal("DB_USER, DB_HOST, DB_NAME are required env vars. Railway plugin vars like MYSQL_HOST, MYSQL_DATABASE, or DATABASE_URL are supported.")
	}

	if port == "" {
		port = "3306"
	}
	if sslMode == "" || sslMode == "true" {
		sslMode = "skip-verify"
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=%s",
		user, pass, host, port, name, sslMode,
	)

	var err error
	// Retry — Railway sometimes needs a few seconds on cold start
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
