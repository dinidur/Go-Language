package handler

import (
	"context"
	"student-microservices/proto"
	"student-microservices/student-service/database"
	"student-microservices/student-service/models"
)

type StudentHandler struct {
	proto.UnimplementedStudentServiceServer
}

// GetStudent — called by Grade Service via gRPC
func (h *StudentHandler) GetStudent(ctx context.Context, req *proto.StudentRequest) (*proto.StudentResponse, error) {
	var student models.Student

	result := database.DB.First(&student, req.Id)
	if result.Error != nil {
		return &proto.StudentResponse{Found: false}, nil
	}

	return &proto.StudentResponse{
		Id:      int32(student.ID),
		Name:    student.Name,
		Subject: student.Subject,
		Grade:   student.Grade,
		Found:   true,
	}, nil
}

// CreateStudent — called by Grade Service via gRPC
func (h *StudentHandler) CreateStudent(ctx context.Context, req *proto.CreateStudentRequest) (*proto.CreateStudentResponse, error) {
	student := models.Student{
		Name:    req.Name,
		Subject: req.Subject,
		Grade:   req.Grade,
	}

	database.DB.Create(&student)

	return &proto.CreateStudentResponse{
		Id:      int32(student.ID),
		Name:    student.Name,
		Subject: student.Subject,
		Grade:   student.Grade,
	}, nil
}