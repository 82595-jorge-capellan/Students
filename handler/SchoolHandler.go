package SchoolHandler

import (
	mapper "github.com/82595-jorge-capellan/mapper"
	pb "github.com/82595-jorge-capellan/protobuf"
	service "github.com/82595-jorge-capellan/service"
)

func AddStudent(in *pb.StudentRequest) (*pb.StudentResponse, error) {

	// Estos 2 pasos se hacen asi porque es preferible pasar proto->json->model que proto->model
	//conversion de proto a json
	jsonStudent := mapper.ProtoStudentToJson(in)
	// transformar estudiante json a estudiante model
	modelStudent := mapper.JsonStudentToModel(jsonStudent)

	//agregamos el estudiante sin un _id especifico en el documento de opensearch
	res, _ := service.AddStudent(modelStudent, "")
	return &pb.StudentResponse{
		Status:     res,
		FinalScore: 0,
	}, nil
}

func AddScoreOfStudent(in *pb.StudentScoreRequest) (*pb.StudentResponse, error) {

	//obtenemos el estudiante con el id que buscamos y el _id del documento de opensearch
	student, docid, _ := service.SearchStudentByID(in.Id)

	switch in.GetExam() {
	case 1:
		student.FirstExam = in.GetScore()
	case 2:
		student.SecondExam = in.GetScore()
	case 3:
		student.ThirdExam = in.GetScore()
	}

	//agregamos el estudiante con un _id de documento especifico para que sobreescriba el estudiante anterior(el mismo)
	res, _ := service.AddStudent(student, docid)
	return &pb.StudentResponse{
		Status:     res,
		FinalScore: 0,
	}, nil
}

func CalculateFinalScore(in *pb.StudentFinalScoreRequest) (*pb.StudentResponse, error) {

	//obtenemos el estudiante con el id que buscamos y el _id del documento de opensearch
	student, docid, _ := service.SearchStudentByID(in.Id)

	exam1 := student.FirstExam
	exam2 := student.SecondExam
	exam3 := student.ThirdExam
	asignments := student.AsignmentScore

	finalScore := (exam1 + exam2 + exam3 + asignments) / 4

	student.FinalScore = finalScore

	//agregamos el estudiante con un _id de documento especifico para que sobreescriba el estudiante anterior(el mismo)
	res, _ := service.AddStudent(student, docid)
	return &pb.StudentResponse{
		Status:     res,
		FinalScore: 0,
	}, nil
}
