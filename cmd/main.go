package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hmailyan/go_ecommerce/internal/app"
	"github.com/hmailyan/go_ecommerce/internal/app/http/middleware"
	"github.com/hmailyan/go_ecommerce/internal/shared/cache"
	"github.com/hmailyan/go_ecommerce/internal/shared/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	db, err := database.New(database.Config{
		DSN:             os.Getenv("DATABASE_DSN"),
		MaxOpenConns:    25,
		MaxIdleConns:    10,
		ConnMaxLifetime: time.Hour,
	})

	rds := cache.NewShardRedis(5, os.Getenv("REDIS_HOST"))
	if err != nil {
		log.Fatalf("DB init failed: %v", err)
	}

	if os.Getenv("DB_AUTOMIGRATE") == "true" {
		database.AutoMigrate(db)
	}

	r := app.BuildRouter(db, rds)

	r.Use(middleware.TimeoutMiddleware(5 * time.Second))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Println("🚀 API running on :8080")
	r.Run(":8080")
}
