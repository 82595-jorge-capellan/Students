package SchoolGateway

import (
	"context"

	handler "github.com/82595-jorge-capellan/handler"
	pb "github.com/82595-jorge-capellan/protobuf"
)

type Server struct {
	pb.UnimplementedSchoolServer
	Handler *handler.Handler
}

func NewServer(h *handler.Handler) *Server {
	return &Server{Handler: h}
}

func (s *Server) AddStudent(_ context.Context, in *pb.StudentRequest) (*pb.StudentResponse, error) {
	return s.Handler.AddStudent(in)
}

func (s *Server) AddScoreOfStudent(_ context.Context, in *pb.StudentScoreRequest) (*pb.StudentResponse, error) {
	return s.Handler.AddScoreOfStudent(in)
}

func (s *Server) CalculateFinalScore(_ context.Context, in *pb.StudentFinalScoreRequest) (*pb.StudentResponse, error) {
	return s.Handler.CalculateFinalScore(in)
}

func (s *Server) SearchStudentByID(_ context.Context, in *pb.StudentSearchRequest) (*pb.StudentSearchResponse, error) {
	return s.Handler.SearchStudentByID(in)
}

func (s *Server) SearchStudentByIDSec(_ context.Context, in *pb.StudentSearchRequest) (*pb.StudentSearchResponse, error) {
	return s.Handler.SearchStudentByIDSec(in)
}

func (s *Server) SearchStudentByIDGo(_ context.Context, in *pb.StudentSearchRequest) (*pb.StudentSearchResponse, error) {
	return s.Handler.SearchStudentByIDGo(in)
}

func (s *Server) SearchStudentByIDMS(_ context.Context, in *pb.StudentSearchRequest) (*pb.StudentSearchResponse, error) {
	return s.Handler.SearchStudentByIDMS(in)
}
