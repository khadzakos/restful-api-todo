package main

import (
	"log"
	"os"

	"main/database"
	"main/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	router := gin.Default()

	router.POST("/api/v1/tasks", handlers.CreateTask)
	router.GET("/api/v1/tasks", handlers.GetAllTasks)
	router.GET("/api/v1/tasks/:id", handlers.GetTask)
	router.PUT("/api/v1/tasks/:id", handlers.UpdateTask)
	router.DELETE("/api/v1/tasks/:id", handlers.DeleteTask)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
