package main

import (
	"student-grade-tracker/database"
	"student-grade-tracker/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	database.ConnectDB()

	// Auto migrate — creates the table automatically
	database.DB.AutoMigrate(&models.Student{})

	// Setup Gin router
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is running!",
		})
	})

	// Routes
	r.GET("/students", controllers.GetStudents)
	r.GET("/students/:id", controllers.GetStudent)
	r.POST("/students", controllers.CreateStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)

	r.Run(":8080")
}
