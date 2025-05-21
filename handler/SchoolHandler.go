package SchoolHandler

import (
	"log"
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	pb "github.com/82595-jorge-capellan/protobuf"
	service "github.com/82595-jorge-capellan/service"
)

type StudentRequestModel struct {
	Id					int32 		`json:"id"`
	FirstName			string		`json:"FirstName"`
	LastName			string		`json:"LastName"`
	FirstExam			int32		`json:"FirstExam"`
	SecondExam			int32		`json:"SecondExam"`
	ThirdExam			int32		`json:"ThirdExam"`
	AsignmentScore		int32		`json:"AsignmentScore"`
	FinalScore			int32		`json:"FinalScore, omitempty"`
}

func AddStudent(in *pb.StudentRequest) (*pb.StudentResponse, error) {

	//conversion de proto a json
	jsonRequest, err := protojson.Marshal(in)
	if err != nil {
        log.Fatalf("Error al convertir proto a JSON: %v", err)
    }

	//recibir map de json de estudiantes del service
	jsonBin, _ := service.GetJSON()

	//apendar el request al json total
	var jsonStudent map[string]interface{}
	err = json.Unmarshal([]byte(jsonRequest), &jsonStudent)
	if err != nil {
    	panic(err)
	}

	jsonBin = append(jsonBin, jsonStudent)

	// //convertir nuevamente a json
	// jsonFinal, err := json.Marshal(jsonBin)
	// if err != nil {
	// 	panic(err)
	// }

	// //enviar el json final a service para updatear
	// response, err := service.UpdateBin(jsonFinal)
	// if err != nil {
	// 	panic(err)
	// }

	return &pb.StudentResponse{
		Status: "temporal",
		// Status: string(response),
		FinalScore: 0,
		}, nil
}

func AddScoreOfStudent(in *pb.StudentScoreRequest) (*pb.StudentResponse, error) {

	//recibir map de json de estudiantes del service
	jsonBin, _ := service.GetJSON()

	for i, student := range jsonBin {
		log.Printf("student %v\n", i)

		if id, ok := student["id"].(float64); ok {
			if int32(id) == in.GetId() {
				if in.GetExam() == 1{
					student["firstExam"] = in.GetScore()
				} else if in.GetExam() == 2 {
					student["secondExam"] = in.GetScore()
				} else {
					student["thirdExam"] = in.GetScore()
				}
			}
		}
	}

	//convertir nuevamente a json
	jsonFinal, err := json.Marshal(jsonBin)
	if err != nil {
		panic(err)
	}

	//enviar el json final a service para updatear
	response, err := service.UpdateBin(jsonFinal)
	if err != nil {
		panic(err)
	}

	return &pb.StudentResponse{
		Status: string(response),
		FinalScore: 0,
		}, nil
}

func CalculateFinalScore(in *pb.StudentFinalScoreRequest) (*pb.StudentResponse, error) {

	//recibir map de json de estudiantes del service
	jsonBin, _ := service.GetJSON()

	for i, student := range jsonBin {
		log.Printf("student %v\n", i)

		if id, ok := student["id"].(float64); ok {
			if int32(id) == in.GetId() {
				exam1 := student["firstExam"].(float64)
				exam2:= student["secondExam"].(float64)
				exam3:= student["thirdExam"].(float64)
				asignments := student["asignmentScore"].(float64)

				finalScore := (exam1 + exam2 + exam3 + asignments)/4

				student["finalScore"] = finalScore

			}
		}
	}

	//convertir nuevamente a json
	jsonFinal, err := json.Marshal(jsonBin)
	if err != nil {
		panic(err)
	}

	//enviar el json final a service para updatear
	response, err := service.UpdateBin(jsonFinal)
	if err != nil {
		panic(err)
	}

	return &pb.StudentResponse{
		Status: string(response),
		FinalScore: 0,
		}, nil
	}




