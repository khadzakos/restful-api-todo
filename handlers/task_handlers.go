package handlers

import (
	"log"
	"main/database"
	"main/database/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.GetDBConnection().Create(&task).Error; err != nil {
		log.Printf("Failed to create task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": task})
	log.Println("Task created")
}

func GetAllTasks(c *gin.Context) {
	tasks := database.GetAllTasks()
	c.JSON(http.StatusOK, gin.H{"data": tasks})
	log.Println("Tasks showed")
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	task := database.GetTask(id)
	c.JSON(http.StatusOK, gin.H{"data": task})
	log.Println("Task showed")
}

func UpdateTask(c *gin.Context) {
	var task models.UpdateTaskModel
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.UpdateTask(&task)
	c.JSON(http.StatusOK, gin.H{"data": task})
	log.Println("Task updated")
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	database.DeleteTask(id)
	c.JSON(http.StatusOK, gin.H{"data": "Task deleted"})
	log.Println("Task deleted")
}
