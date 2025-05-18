package SchoolHandler

import (
	"google.golang.org/protobuf/encoding/protojson"
	pb "github.com/82595-jorge-capellan/protobuf"
	service "github.com/82595-jorge-capellan/service"
)

func AddStudent(in *pb.StudentRequest) (*pb.StudentResponse, error) {
	json, _ := ProtoToJSON(in)
	return service.AddStudent(json)
}

func ProtoToJSON(msg *pb.StudentRequest) (string, error) {
	marshaler := protojson.MarshalOptions{
		Indent:          "  ", // opcional: agrega formato bonito
		UseProtoNames:   true, // usa nombres del proto (en lugar de camelCase)
		EmitUnpopulated: true, // incluye campos con valor cero
	}

	jsonBytes, err := marshaler.Marshal(msg)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}