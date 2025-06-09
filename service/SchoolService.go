package SchoolService

import (
	"fmt"
	"time"

	mapper "github.com/82595-jorge-capellan/mapper"
	"github.com/82595-jorge-capellan/repo"
)

type Service struct {
	Repo *repo.Repo
}

func NewService(r *repo.Repo) *Service {
	return &Service{Repo: r}
}

func (s *Service) AddStudent(student mapper.StudentRequestModel, optionalID string) (string, error) {
	res, err := s.Repo.AddStudent(&student, optionalID)
	return res, err
}

func (s *Service) SearchStudentByID(id int32, subject string) (mapper.StudentRequestModel, string, error) {

	res, doc_id, err := s.Repo.SearchStudentByID(id, subject)
	return res, doc_id, err
}

func (s *Service) SearchStudentByIDAllSubjects(_ int32) []mapper.StudentRequestModel {

	s.msearch(3)
	s.secuentialSearch(1)
	subroutineSearch := s.subRoutineSearch(2)

	//return resSecuential[0:3]
	//return resMSearch[0:3]
	return subroutineSearch[0:3]
}

func (s *Service) SearchStudentByIDAllSubjectsSec(id int32) []mapper.StudentRequestModel {

	secuentialSearch := s.secuentialSearch(id)

	return secuentialSearch[0:3]
}

func (s *Service) SearchStudentByIDAllSubjectsGo(id int32) []mapper.StudentRequestModel {

	subroutineSearch := s.subRoutineSearch(id)

	return subroutineSearch[0:3]
}

func (s *Service) SearchStudentByIDAllSubjectsMS(id int32) []mapper.StudentRequestModel {

	MSSearch := s.msearch(id)

	return MSSearch[0:3]
}

func (s *Service) msearch(id int32) [3]mapper.StudentRequestModel {
	startTime := time.Now()
	res := s.Repo.MsearchSearchStudent(id)

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("----------------  %v  --------------------\n", id)
	fmt.Printf("Multi Search request lasted %v nanoseconds\n", duration.Nanoseconds())
	return [3]mapper.StudentRequestModel(res)
}

func (s *Service) secuentialSearch(id int32) [3]mapper.StudentRequestModel {

	startTime := time.Now()
	resbiology, _, _ := s.Repo.SearchStudentByID(id, "biology")
	reschemistry, _, _ := s.Repo.SearchStudentByID(id, "chemistry")
	resmath, _, _ := s.Repo.SearchStudentByID(id, "math")

	var res [3]mapper.StudentRequestModel
	res[0] = resbiology
	res[1] = reschemistry
	res[2] = resmath

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("----------------  %v  --------------------\n", id)
	fmt.Printf("secuential Search request lasted %v nanoseconds\n", duration.Nanoseconds())
	return res
}

func (s *Service) subRoutineSearch(id int32) [3]mapper.StudentRequestModel {
	startTime := time.Now()
	res := make(chan mapper.StudentRequestModel, 3)
	var result [3]mapper.StudentRequestModel
	subjects := []string{"biology", "math", "chemistry"}

	for _, subject := range subjects {
		go func(subject string) {
			resbiology, _, _ := s.Repo.SearchStudentByID(id, subject)
			res <- resbiology
		}(subject)
	}
	for i := 0; i < 3; i++ {
		result[i] = <-res
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("----------------  %v  --------------------\n", id)
	fmt.Printf("goroutine Search request lasted %v nanoseconds\n", duration.Nanoseconds())
	return result
}
