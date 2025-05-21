package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/82595-jorge-capellan/protobuf"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSchoolClient(conn)

	// Contact the server and print out its response.
	//AddStudent(c)

	AddScoreOfStudent(c)


}

func AddStudent(c pb.SchoolClient){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.AddStudent(ctx, &pb.StudentRequest{
		Id: 3,
		FirstName: "Federico",
		LastName: "Fabra",
		FirstExam: 2,
		SecondExam: 2,
		ThirdExam: 2,
		AsignmentScore: 2,
		FinalScore: 0,
	})
	if err != nil {
		log.Fatalf("could not add student: %v", err)
	}
	log.Printf("student added: %s", r.GetStatus())
}

func AddScoreOfStudent(c pb.SchoolClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.AddScoreOfStudent(ctx, &pb.StudentScoreRequest{
		Id: 1,
		Exam: 1,
		Score: 10,
	})
	if err != nil {
		log.Fatalf("could not add student: %v", err)
	}
	log.Printf("student added: %s", r.GetStatus())
}