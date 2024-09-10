package database

import (
	"fmt"
	"log"
	"main/database/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	_DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = _DB
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
	log.Println("Database connected")
}

func GetDBConnection() *gorm.DB {
	if DB == nil {
		log.Fatal("Database not initialized")
	}
	return DB
}

func AddTask(t *models.Task) {
	if GetDBConnection().Create(t).Error != nil {
		log.Fatal("Failed to create task")
	}
}

func UpdateTask(t *models.UpdateTaskModel) {
	if GetDBConnection().Model(&models.Task{}).Where("id = ?", t.ID).Updates(t).Error != nil {
		log.Fatal("Failed to update task")
	}
}

func GetAllTasks() []models.Task {
	var tasks []models.Task
	if GetDBConnection().Find(&tasks).Error != nil {
		log.Fatal("Failed to get tasks")
	}
	return tasks
}

func GetTask(id string) models.Task {
	var task models.Task
	if GetDBConnection().First(&task, id).Error != nil {
		log.Fatal("Failed to get task")
	}
	return task
}

func DeleteTask(id string) {
	if GetDBConnection().Delete(&models.Task{}, id).Error != nil {
		log.Fatal("Failed to delete task")
	}
}
