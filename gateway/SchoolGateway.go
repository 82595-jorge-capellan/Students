package schoolGateway

import (

	pb "github.com/82595-jorge-capellan/protobuf"
	handler "github.com/82595-jorge-capellan/handler"
)

// var (
// 	port = flag.Int("port", 50051, "The server port")
// )

// server is used to implement helloworld.GreeterServer.
// type server struct {
// 	pb.UnimplementedDataServer
// }


// SayHello implements helloworld.GreeterServer
func AddStudent(in *pb.StudentRequest) (*pb.StudentResponse, error) {
	return handler.AddStudent(in)
}

// func main() {
// 	flag.Parse()
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterDataServer(s, &server{})
// 	log.Printf("server listening at %v", lis.Addr())
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }