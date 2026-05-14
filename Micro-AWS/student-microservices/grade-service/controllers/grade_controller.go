package controllers

import (
	"context"
	"net/http"
	"student-microservices/grade-service/database"
	"student-microservices/grade-service/models"
	"student-microservices/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// connectToStudentService — connects to Student Service via gRPC
func connectToStudentService() (proto.StudentServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient("localhost:9090",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, err
	}
	client := proto.NewStudentServiceClient(conn)
	return client, conn, nil
}

// AddGrade — adds a grade and verifies student exists via gRPC
func AddGrade(c *gin.Context) {
	var grade models.Grade

	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call Student Service via gRPC to verify student exists
	client, conn, err := connectToStudentService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not connect to Student Service"})
		return
	}
	defer conn.Close()

	// Check if student exists
	studentResp, err := client.GetStudent(context.Background(), &proto.StudentRequest{
		Id: int32(grade.StudentID),
	})

	if err != nil || !studentResp.Found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
		return
	}

	// Save grade to database
	database.DB.Create(&grade)

	c.JSON(http.StatusCreated, gin.H{
		"data":    grade,
		"student": studentResp.Name,
	})
}

// GetGrades — get all grades for a student
func GetGrades(c *gin.Context) {
	var grades []models.Grade
	studentID := c.Param("student_id")

	database.DB.Where("student_id = ?", studentID).Find(&grades)

	c.JSON(http.StatusOK, gin.H{
		"data": grades,
	})
}