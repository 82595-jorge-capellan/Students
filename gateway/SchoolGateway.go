package SchoolGateway

import (
	"context"
	pb "github.com/82595-jorge-capellan/protobuf"
	handler "github.com/82595-jorge-capellan/handler"
)

type Server struct {
	pb.UnimplementedSchoolServer
}


func (s *Server) AddStudent(_ context.Context, in *pb.StudentRequest) (*pb.StudentResponse, error) {
	return handler.AddStudent(in)
}
