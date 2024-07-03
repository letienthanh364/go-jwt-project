package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	routes "github.com/teddylethal/golang-jwt-project/routes"
	"log"
	"os"
)

func main() {
	errLoadEnv := godotenv.Load(".env")
	if errLoadEnv != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api version 1"})
	})

	router.GET("/v2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api version 2"})
	})

	router.Run(":" + port)
}
