package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type Student struct {
	Id             int32   `json:"id"`
	FirstName      string  `json:"FirstName"`
	LastName       string  `json:"LastName"`
	FirstExam      int32   `json:"FirstExam"`
	SecondExam     int32   `json:"SecondExam"`
	ThirdExam      int32   `json:"ThirdExam"`
	AsignmentScore int32   `json:"AsignmentScore"`
	FinalScore     float32 `json:"FinalScore,omitempty"`
	Subject        string  `json:"Subject"`
}

func main() {
	const subject = "chemistry"

	const num = 1000

	file, _ := os.Create("bulk_create_chemistry.ndjson")

	defer file.Close()

	for i := 1; i <= num; i++ {
		firstExam := rand.Int31n(11)
		secondExam := rand.Int31n(11)
		thirdExam := rand.Int31n(11)
		asignment := rand.Int31n(11)

		final := float32(firstExam+secondExam+thirdExam+asignment) / 4.0

		student := Student{
			Id:             int32(i),
			FirstName:      fmt.Sprintf("Nombre%d", i),
			LastName:       fmt.Sprintf("Apellido%d", i),
			FirstExam:      firstExam,
			SecondExam:     secondExam,
			ThirdExam:      thirdExam,
			AsignmentScore: asignment,
			FinalScore:     final,
			Subject:        subject,
		}

		// Línea de acción bulk (usa índice en URL, así que no se especifica aquí)
		meta := map[string]any{"index": map[string]any{}}
		metaJSON, _ := json.Marshal(meta)
		docJSON, _ := json.Marshal(student)

		// Escribe ambas líneas (meta y doc)
		file.Write(metaJSON)
		file.Write([]byte("\n"))
		file.Write(docJSON)
		file.Write([]byte("\n"))
	}

	fmt.Println("Archivo NDJSON generado: bulk_create.ndjson")
}
