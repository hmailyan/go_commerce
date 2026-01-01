package gocommerce

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hmailyan/go_ecommerce/controllers"
	"github.com/hmailyan/go_ecommerce/database"
	"github.com/hmailyan/go_ecommerce/middleware"
	"github.com/hmailyan/go_ecommerce/routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveFromCart())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
