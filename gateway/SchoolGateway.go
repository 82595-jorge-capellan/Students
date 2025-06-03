package SchoolGateway

import (
	"context"

	handler "github.com/82595-jorge-capellan/handler"
	pb "github.com/82595-jorge-capellan/protobuf"
)

type Server struct {
	pb.UnimplementedSchoolServer
}

func (s *Server) AddStudent(_ context.Context, in *pb.StudentRequest) (*pb.StudentResponse, error) {
	return handler.AddStudent(in)
}

func (s *Server) AddScoreOfStudent(_ context.Context, in *pb.StudentScoreRequest) (*pb.StudentResponse, error) {
	return handler.AddScoreOfStudent(in)
}

func (s *Server) CalculateFinalScore(_ context.Context, in *pb.StudentFinalScoreRequest) (*pb.StudentResponse, error) {
	return handler.CalculateFinalScore(in)
}

func (s *Server) SearchStudentByID(_ context.Context, in *pb.StudentSearchRequest) (*pb.StudentSearchResponse, error) {
	return handler.SearchStudentByID(in)
}
