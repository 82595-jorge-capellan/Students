package schoolMapper

import (
	"encoding/json"
	"log"

	pb "github.com/82595-jorge-capellan/protobuf"
	"google.golang.org/protobuf/encoding/protojson"
)

type StudentRequestModel struct {
	Id             int32  `json:"id"`
	FirstName      string `json:"FirstName"`
	LastName       string `json:"LastName"`
	FirstExam      int32  `json:"FirstExam"`
	SecondExam     int32  `json:"SecondExam"`
	ThirdExam      int32  `json:"ThirdExam"`
	AsignmentScore int32  `json:"AsignmentScore"`
	FinalScore     int32  `json:"FinalScore,omitempty"`
}

func ProtoStudentToJson(proto *pb.StudentRequest) []byte {
	json, err := protojson.Marshal(proto)
	if err != nil {
		log.Fatalf("Error al convertir proto a JSON: %v", err)
	}
	return json
}

func JsonStudentToModel(j []byte) StudentRequestModel {
	var model StudentRequestModel
	err := json.Unmarshal(j, &model)
	if err != nil {
		panic(err)
	}
	return model
}

func JsonToStudentArray(j []byte) []StudentRequestModel {
	var modelArray []StudentRequestModel
	err := json.Unmarshal(j, &modelArray)
	if err != nil {
		panic(err)
	}
	return modelArray
}

func StudentArrayToJson(s []StudentRequestModel) []byte {
	j, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return j
}
