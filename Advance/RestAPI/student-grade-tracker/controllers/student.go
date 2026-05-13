package controllers

import (
	"net/http"
	"student-grade-tracker/database"
	"student-grade-tracker/models"

	"github.com/gin-gonic/gin"
)

// Get all students
func GetStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, gin.H{
		"data": students,
	})
}

// Get single student
func GetStudent(c *gin.Context) {
	var student models.Student
	if err := database.DB.First(&student, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}

// Create student
func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusCreated, gin.H{
		"data": student,
	})
}

// Update student
func UpdateStudent(c *gin.Context) {
	var student models.Student
	if err := database.DB.First(&student, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found!",
		})
		return
	}
	var input models.Student
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&student).Updates(input)
	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}

// Delete student
func DeleteStudent(c *gin.Context) {
	var student models.Student
	if err := database.DB.First(&student, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found!",
		})
		return
	}
	database.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted successfully!",
	})
}