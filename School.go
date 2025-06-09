package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	gateway "github.com/82595-jorge-capellan/gateway"
	handler "github.com/82595-jorge-capellan/handler"
	pb "github.com/82595-jorge-capellan/protobuf"
	"github.com/82595-jorge-capellan/repo"
	service "github.com/82595-jorge-capellan/service"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {

	r, err := repo.NewRepo()
	if err != nil {
		log.Fatalf("error creando repo: %v", err)
	}

	svc := service.NewService(r)
	h := handler.NewHandler(svc)
	server := gateway.NewServer(h)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSchoolServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
