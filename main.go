package main

import (
	"jwttuts/controllers"
	"jwttuts/database"
	"jwttuts/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	DB := database.Connect()

	boil.SetDB(DB)
	router := initRouter()
	router.Run(":8080")

}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}

	}

	return router

}
