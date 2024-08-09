package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	//allow only localhost:8080 as origin to prevent CORS issues
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://foo.com", "http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin"},
	}))

	server.Run(":8080")

}
