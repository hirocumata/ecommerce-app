package main

import (
	
	"github.com/hirocumata/ecommerce-app/controllers"
	"github.com/hirocumata/ecommerce-app/database"
	"github.com/hirocumata/ecommerce-app/middleware"
	"github.com/hirocumata/ecommerce-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os - Getenv("PORT")
	if port == "" {
		port = 8080
	}

	app := controllers.NewApplication(database.ProductData(database.client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(Gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
