package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"context"

	"google.golang.org/grpc"
	pb "github.com/82595-jorge-capellan/protobuf"
	gateway "github.com/82595-jorge-capellan/gateway"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedSchoolServer
}

func (s *server) AddStudent(_ context.Context, in *pb.StudentRequest) (*pb.StudentResponse, error) {
	return gateway.AddStudent(in)
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSchoolServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}