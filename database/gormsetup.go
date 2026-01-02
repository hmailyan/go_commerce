package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hmailyan/go_ecommerce/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupGORM() *gorm.DB {
	if DB != nil {
		return DB
	}

	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		host := os.Getenv("POSTGRES_HOST")
		if host == "" {
			host = "localhost"
		}
		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		dbname := os.Getenv("POSTGRES_DB")
		port := os.Getenv("POSTGRES_PORT")
		if port == "" {
			port = "5432"
		}
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, dbname, port)
	}

	var db *gorm.DB
	var err error
	// Retry loop to wait for Postgres to come up (useful when starting containers)
	for i := 0; i < 8; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("failed to connect to Postgres (attempt %d): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("failed to connect to Postgres: %v", err)
	}

	DB = db

	// AutoMigrate models
	if err := DB.AutoMigrate(&models.User{}, &models.Product{}, &models.ProductUser{}, &models.Address{}, &models.Order{}); err != nil {
		log.Fatalf("failed to automigrate models: %v", err)
	}

	return DB
}
