package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hmailyan/go_ecommerce/controllers"
	"github.com/hmailyan/go_ecommerce/database"
	"github.com/hmailyan/go_ecommerce/middleware"
	"github.com/hmailyan/go_ecommerce/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	// Initialize Postgres + GORM
	database.SetupGORM()

	app := controllers.NewApplication()

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	// Cart endpoints (require authentication)
	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveFromCart())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	// Address endpoints (protected)
	router.POST("/users/address", controllers.AddAddress())
	router.PUT("/users/address/home", controllers.EditHomeAddress())
	router.PUT("/users/address/work", controllers.EditWorkAddress())
	router.DELETE("/users/address", controllers.DeleteAddress())

	log.Fatal(router.Run(":" + port))
}
