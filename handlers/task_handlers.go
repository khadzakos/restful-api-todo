package handlers

import (
	"encoding/json"
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

	beautifiedJSON, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to beautify JSON"})
		return
	}

	c.Data(http.StatusCreated, "application/json", beautifiedJSON)

	log.Println("Task created")
}

func GetAllTasks(c *gin.Context) {
	tasks := database.GetAllTasks()

	beautifiedJSON, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to beautify JSON"})
		return
	}

	c.Data(http.StatusOK, "application/json", beautifiedJSON)

	// c.JSON(http.StatusOK, gin.H{"data": tasks})
	log.Println("Tasks showed")
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	task := database.GetTask(id)

	beautifiedJSON, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to beautify JSON"})
		return
	}

	c.Data(http.StatusOK, "application/json", beautifiedJSON)

	// c.JSON(http.StatusOK, gin.H{"data": task})
	log.Println("Task showed")
}

func UpdateTask(c *gin.Context) {
	var task models.UpdateTaskModel
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.UpdateTask(&task)

	beautifiedJSON, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to beautify JSON"})
		return
	}

	c.Data(http.StatusOK, "application/json", beautifiedJSON)
	log.Println("Task updated")
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	database.DeleteTask(id)
	c.JSON(http.StatusOK, gin.H{"data": "Task deleted"})
	log.Println("Task deleted")
}
