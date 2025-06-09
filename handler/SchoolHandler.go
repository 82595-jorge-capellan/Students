package SchoolHandler

import (
	"math"

	mapper "github.com/82595-jorge-capellan/mapper"
	pb "github.com/82595-jorge-capellan/protobuf"
	service "github.com/82595-jorge-capellan/service"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) AddStudent(in *pb.StudentRequest) (*pb.StudentResponse, error) {

	// Estos 2 pasos se hacen asi porque es preferible pasar proto->json->model que proto->model
	//conversion de proto a json
	jsonStudent := mapper.ProtoStudentToJson(in)
	// transformar estudiante json a estudiante model
	modelStudent := mapper.JsonStudentToModel(jsonStudent)

	//agregamos el estudiante sin un _id especifico en el documento de opensearch
	res, _ := h.Service.AddStudent(modelStudent, "")
	return &pb.StudentResponse{
		Status:     res,
		FinalScore: 0,
	}, nil
}

func (h *Handler) AddScoreOfStudent(in *pb.StudentScoreRequest) (*pb.StudentResponse, error) {

	//obtenemos el estudiante con el id que buscamos y el _id del documento de opensearch
	student, docid, _ := h.Service.SearchStudentByID(in.Id, in.Subject)

	switch in.GetExam() {
	case 1:
		student.FirstExam = in.GetScore()
	case 2:
		student.SecondExam = in.GetScore()
	case 3:
		student.ThirdExam = in.GetScore()
	}

	//agregamos el estudiante con un _id de documento especifico para que sobreescriba el estudiante anterior(el mismo)
	res, _ := h.Service.AddStudent(student, docid)
	return &pb.StudentResponse{
		Status:     res,
		FinalScore: 0,
	}, nil
}

func (h *Handler) CalculateFinalScore(in *pb.StudentFinalScoreRequest) (*pb.StudentResponse, error) {

	//obtenemos el estudiante con el id que buscamos y el _id del documento de opensearch
	student, docid, _ := h.Service.SearchStudentByID(in.Id, in.Subject)

	exam1 := student.FirstExam
	exam2 := student.SecondExam
	exam3 := student.ThirdExam
	asignments := student.AsignmentScore

	sum := float32(exam1 + exam2 + exam3 + asignments)
	finalScore := sum / 4
	roundedScore := float32(math.Round(float64(finalScore)*100) / 100)

	student.FinalScore = roundedScore

	//agregamos el estudiante con un _id de documento especifico para que sobreescriba el estudiante anterior(el mismo)
	res, _ := h.Service.AddStudent(student, docid)
	return &pb.StudentResponse{
		Status:     res,
		FinalScore: 0,
	}, nil
}

func (h *Handler) SearchStudentByID(in *pb.StudentSearchRequest) (*pb.StudentSearchResponse, error) {

	students := h.Service.SearchStudentByIDAllSubjects(in.Id)

	var grpcStudents []*pb.StudentRequest

	for _, s := range students {
		grpcStudent := &pb.StudentRequest{
			Id:             s.Id,
			FirstName:      s.FirstName,
			LastName:       s.LastName,
			FirstExam:      s.FirstExam,
			SecondExam:     s.SecondExam,
			ThirdExam:      s.ThirdExam,
			AsignmentScore: s.AsignmentScore,
			FinalScore:     s.FinalScore,
			Subject:        s.Subject,
		}
		grpcStudents = append(grpcStudents, grpcStudent)
	}

	return &pb.StudentSearchResponse{
		StudentSubject: grpcStudents,
	}, nil
}

func (h *Handler) SearchStudentByIDSec(in *pb.StudentSearchRequest) (*pb.StudentSearchResponse, error) {

	students := h.Service.SearchStudentByIDAllSubjectsSec(in.Id)

	var grpcStudents []*pb.StudentRequest

	for _, s := range students {
		grpcStudent := &pb.StudentRequest{
			Id:             s.Id,
			FirstName:      s.FirstName,
			LastName:       s.LastName,
			FirstExam:      s.FirstExam,
			SecondExam:     s.SecondExam,
			ThirdExam:      s.ThirdExam,
			AsignmentScore: s.AsignmentScore,
			FinalScore:     s.FinalScore,
			Subject:        s.Subject,
		}
		grpcStudents = append(grpcStudents, grpcStudent)
	}

	return &pb.StudentSearchResponse{
		StudentSubject: grpcStudents,
	}, nil
}

func (h *Handler) SearchStudentByIDGo(in *pb.StudentSearchRequest) (*pb.StudentSearchResponse, error) {

	students := h.Service.SearchStudentByIDAllSubjectsGo(in.Id)

	var grpcStudents []*pb.StudentRequest

	for _, s := range students {
		grpcStudent := &pb.StudentRequest{
			Id:             s.Id,
			FirstName:      s.FirstName,
			LastName:       s.LastName,
			FirstExam:      s.FirstExam,
			SecondExam:     s.SecondExam,
			ThirdExam:      s.ThirdExam,
			AsignmentScore: s.AsignmentScore,
			FinalScore:     s.FinalScore,
			Subject:        s.Subject,
		}
		grpcStudents = append(grpcStudents, grpcStudent)
	}

	return &pb.StudentSearchResponse{
		StudentSubject: grpcStudents,
	}, nil
}

func (h *Handler) SearchStudentByIDMS(in *pb.StudentSearchRequest) (*pb.StudentSearchResponse, error) {

	students := h.Service.SearchStudentByIDAllSubjectsMS(in.Id)

	var grpcStudents []*pb.StudentRequest

	for _, s := range students {
		grpcStudent := &pb.StudentRequest{
			Id:             s.Id,
			FirstName:      s.FirstName,
			LastName:       s.LastName,
			FirstExam:      s.FirstExam,
			SecondExam:     s.SecondExam,
			ThirdExam:      s.ThirdExam,
			AsignmentScore: s.AsignmentScore,
			FinalScore:     s.FinalScore,
			Subject:        s.Subject,
		}
		grpcStudents = append(grpcStudents, grpcStudent)
	}

	return &pb.StudentSearchResponse{
		StudentSubject: grpcStudents,
	}, nil
}
