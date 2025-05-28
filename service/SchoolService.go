package SchoolService

import (
	mapper "github.com/82595-jorge-capellan/mapper"
	"github.com/82595-jorge-capellan/repo"
)

func AddStudent(student mapper.StudentRequestModel, optionalID string) (string, error) {
	res, err := repo.AddStudent(&student, optionalID)
	return res, err
}

func SearchStudentByID(id int32) (mapper.StudentRequestModel, string, error) {

	res, doc_id, err := repo.SearchStudentByID(id)
	return res, doc_id, err
}
