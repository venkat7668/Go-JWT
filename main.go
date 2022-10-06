package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/venkat7668/Go-JWT/controllers"
	"github.com/venkat7668/Go-JWT/database"
	"github.com/venkat7668/Go-JWT/middlewares"
	_ "gorm.io/driver/mysql"
)

func main() {
	// Initialize Database
	DB_HOST := os.Getenv("MYSQL_HOST")
	DB_PORT := os.Getenv("MYSQL_PORT")
	DB_NAME := os.Getenv("MYSQL_DB")
	DB_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	DB_USER := os.Getenv("MYSQL_USER")
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	// "root:helloworld@tcp(localhost:3308)/testapp?parseTime=true"
	database.Connect(connectionStr)
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
