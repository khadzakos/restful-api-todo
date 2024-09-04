package handlers

import (
	"main/database"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	c.BindJSON(&task)

	database.DB.Create(&task)

	c.JSON(http.StatusCreated, task)
}

func GetTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")
	database.DB.First(&task, id)

	c.JSON(http.StatusOK, task)
}

func GetAllTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")
	database.DB.First(&task, id)

	c.BindJSON(&task)

	database.DB.Save(&task)

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")
	database.DB.First(&task, id)

	database.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"id" + id: "deleted"})
}
