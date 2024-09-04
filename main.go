package main

import (
	"log"

	"main/database"
	"main/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	router := gin.Default()

	router.POST("/api/tasks", handlers.CreateTask)
	router.GET("/api/tasks", handlers.GetAllTasks)
	router.GET("/api/tasks/{id}", handlers.GetTask)
	router.PUT("/api/tasks/{id}", handlers.UpdateTask)
	router.DELETE("api/tasks/{id}", handlers.DeleteTask)

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
