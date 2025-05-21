package main

import (
	"context"
	"flag"
	"log"
	"time"
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"

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

	//AddScoreOfStudent(c)

	//CalculateFinalScore(c)



	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("=== Menú Principal ===")
		fmt.Println("1. Agregar estudiante")
		fmt.Println("2. Cambiar nota de estudiante")
		fmt.Println("3. Calcular nota final de estudiante")
		fmt.Println("4. Salir")
		fmt.Print("Elija una opción: ")

		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			AddStudent(reader, c)
		case "2":
			AddScoreOfStudent(reader, c)
		case "3":
			CalculateFinalScore(reader, c)
		case "4":
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción no válida. Intente nuevamente.")
		}

		fmt.Println()
	}


}

func AddStudent(reader *bufio.Reader, c pb.SchoolClient) {
	fmt.Println("\n--- Ingreso de datos de persona ---")
	fmt.Print("Id: ")
	idString, _ := reader.ReadString('\n')
	idString = strings.TrimSpace(idString)
	id64, _ := strconv.ParseInt(idString, 10, 32)
	id := int32(id64)

	fmt.Printf("valor %s", id)


	fmt.Print("Nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	fmt.Print("Apellido: ")
	apellido, _ := reader.ReadString('\n')
	apellido = strings.TrimSpace(apellido)

	fmt.Print("primer examen: ")
	primerExamenString, _ := reader.ReadString('\n')
	primerExamenString = strings.TrimSpace(primerExamenString)
	primerExamen64, _ := strconv.ParseInt(primerExamenString, 10, 32)
	primerExamen := int32(primerExamen64)

	fmt.Print("segundo examen: ")
	segundoExamenString, _ := reader.ReadString('\n')
	segundoExamenString = strings.TrimSpace(segundoExamenString)
	segundoExamen64, _ := strconv.ParseInt(segundoExamenString, 10, 32)
	segundoExamen := int32(segundoExamen64)

	fmt.Print("tercer examen: ")
	tercerExamenString, _ := reader.ReadString('\n')
	tercerExamenString = strings.TrimSpace(tercerExamenString)
	tercerExamen64, _ := strconv.ParseInt(tercerExamenString, 10, 32)
	tercerExamen := int32(tercerExamen64)

	fmt.Print("trabajos practicos: ")
	trabajosPracticosString, _ := reader.ReadString('\n')
	trabajosPracticosString = strings.TrimSpace(trabajosPracticosString)
	trabajosPracticos64, _ := strconv.ParseInt(trabajosPracticosString, 10, 32)
	trabajosPracticos := int32(trabajosPracticos64)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.AddStudent(ctx, &pb.StudentRequest{
		Id: id,
		FirstName: nombre,
		LastName: apellido,
		FirstExam: primerExamen,
		SecondExam: segundoExamen,
		ThirdExam: tercerExamen,
		AsignmentScore: trabajosPracticos,
		FinalScore: 0,
	})
	if err != nil {
		log.Fatalf("could not add student: %v", err)
	}
	log.Printf("student added: %s", r.GetStatus())
}


func AddScoreOfStudent(reader *bufio.Reader, c pb.SchoolClient) {
	fmt.Println("\n--- Ingreso de datos de persona ---")
	fmt.Print("Id: ")
	idString, _ := reader.ReadString('\n')
	idString = strings.TrimSpace(idString)
	id64, _ := strconv.ParseInt(idString, 10, 32)
	id := int32(id64)

	fmt.Print("exam (1 - 2 - 3): ")
	examString, _ := reader.ReadString('\n')
	examString = strings.TrimSpace(examString)
	exam64, _ := strconv.ParseInt(examString, 10, 32)
	exam := int32(exam64)

	fmt.Print("Score: ")
	scoreString, _ := reader.ReadString('\n')
	scoreString = strings.TrimSpace(scoreString)
	score64, _ := strconv.ParseInt(scoreString, 10, 32)
	score := int32(score64)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.AddScoreOfStudent(ctx, &pb.StudentScoreRequest{
		Id: id,
		Exam: exam,
		Score: score,
	})
	if err != nil {
		log.Fatalf("could not add student: %v", err)
	}
	log.Printf("Score changed: %s", r.GetStatus())
}

func CalculateFinalScore(reader *bufio.Reader, c pb.SchoolClient) {

	fmt.Println("\n--- Ingreso de datos de persona ---")
	fmt.Print("Id: ")
	idString, _ := reader.ReadString('\n')
	idString = strings.TrimSpace(idString)
	id64, _ := strconv.ParseInt(idString, 10, 32)
	id := int32(id64)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.CalculateFinalScore(ctx, &pb.StudentFinalScoreRequest{
		Id: id,
	})
	if err != nil {
		log.Fatalf("could not add student: %v", err)
	}
	log.Printf("student added: %s", r.GetStatus())
}