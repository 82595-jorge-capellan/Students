package SchoolService

import (
	"fmt"
	"time"

	mapper "github.com/82595-jorge-capellan/mapper"
	"github.com/82595-jorge-capellan/repo"
)

func AddStudent(student mapper.StudentRequestModel, optionalID string) (string, error) {
	res, err := repo.AddStudent(&student, optionalID)
	return res, err
}

func SearchStudentByID(id int32, subject string) (mapper.StudentRequestModel, string, error) {

	res, doc_id, err := repo.SearchStudentByID(id, subject)
	return res, doc_id, err
}

func SearchStudentByIDAllSubjects(_ int32) []mapper.StudentRequestModel {

	msearch(3)
	secuentialSearch(1)
	subroutineSearch := subRoutineSearch(2)

	//return resSecuential[0:3]
	//return resMSearch[0:3]
	return subroutineSearch[0:3]
}

func msearch(id int32) [3]mapper.StudentRequestModel {
	startTime := time.Now()
	res := repo.MsearchSearchStudent(id)

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("----------------  %V  --------------------\n", id)
	fmt.Printf("Multi Search request lasted %v nanoseconds\n", duration.Nanoseconds())
	return [3]mapper.StudentRequestModel(res)
}

func secuentialSearch(id int32) [3]mapper.StudentRequestModel {

	startTime := time.Now()
	resbiology, _, _ := repo.SearchStudentByID(id, "biology")
	reschemistry, _, _ := repo.SearchStudentByID(id, "chemistry")
	resmath, _, _ := repo.SearchStudentByID(id, "math")

	var res [3]mapper.StudentRequestModel
	res[0] = resbiology
	res[1] = reschemistry
	res[2] = resmath

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("----------------  %V  --------------------\n", id)
	fmt.Printf("secuential Search request lasted %v nanoseconds\n", duration.Nanoseconds())
	return res
}

func subRoutineSearch(id int32) [3]mapper.StudentRequestModel {
	startTime := time.Now()
	res := make(chan mapper.StudentRequestModel, 3)
	var result [3]mapper.StudentRequestModel
	subjects := []string{"biology", "math", "chemistry"}

	for _, subject := range subjects {
		go func(subject string) {
			resbiology, _, _ := repo.SearchStudentByID(id, subject)
			res <- resbiology
		}(subject)
	}
	for i := 0; i < 3; i++ {
		result[i] = <-res
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("----------------  %V  --------------------\n", id)
	fmt.Printf("goroutine Search request lasted %v nanoseconds\n", duration.Nanoseconds())
	return result
}
