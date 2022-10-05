package main

import (
	"github.com/gin-gonic/gin"
	"github.com/venkat7668/Go-JWT/controllers"
	"github.com/venkat7668/Go-JWT/database"
	"github.com/venkat7668/Go-JWT/middlewares"
	_ "gorm.io/driver/mysql"
)

func main() {
	// Initialize Database
	database.Connect("root:helloworld@tcp(localhost:3308)/testapp?parseTime=true")
	database.Migrate()
	router := initRouter()
	router.Run(":80")
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
