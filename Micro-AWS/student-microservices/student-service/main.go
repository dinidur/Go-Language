package main

import (
	"fmt"
	"log"
	"net"

	"student-microservices/proto"
	"student-microservices/student-service/database"
	"student-microservices/student-service/handler"
	"student-microservices/student-service/models"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// Connect to database
	database.ConnectDB()
	database.DB.AutoMigrate(&models.Student{})

	// Start gRPC server in background
	go func() {
		grpcServer := grpc.NewServer()
		proto.RegisterStudentServiceServer(grpcServer, &handler.StudentHandler{})

		lis, err := net.Listen("tcp", ":9090")
		if err != nil {
			log.Fatal("Failed to listen:", err)
		}

		fmt.Println("Student Service gRPC running on port 9090...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("Failed to serve:", err)
		}
	}()

	// HTTP server for direct access
	r := gin.Default()

	r.POST("/students", func(c *gin.Context) {
		var student models.Student
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		database.DB.Create(&student)
		c.JSON(201, gin.H{"data": student})
	})

	r.GET("/students", func(c *gin.Context) {
		var students []models.Student
		database.DB.Find(&students)
		c.JSON(200, gin.H{"data": students})
	})

	fmt.Println("Student Service HTTP running on port 8080...")
	r.Run(":8080")
}