package main

import (
	"fmt"

	"student-microservices/grade-service/controllers"
	"student-microservices/grade-service/database"
	"student-microservices/grade-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	database.ConnectDB()
	database.DB.AutoMigrate(&models.Grade{})

	// Setup Gin router
	r := gin.Default()

	// Routes
	r.POST("/grades", controllers.AddGrade)
	r.GET("/grades/:student_id", controllers.GetGrades)

	fmt.Println("Grade Service running on port 8081...")
	r.Run(":8081")
}