package repo

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mapper "github.com/82595-jorge-capellan/mapper"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

// Struct usada para tener a mano tanto el _id del documento como el studentModel
// agregando mas campos a este struct podemos mapear 1 a 1 la respuesta de Opensearch
type OpenSearchResponse struct {
	Hits struct {
		Hits []struct {
			ID     string                     `json:"_id"`
			Source mapper.StudentRequestModel `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func AddStudent(student *mapper.StudentRequestModel, optionalID string) (string, error) {

	//creamos el cliente de opensearch con ip, credenciales y configuraciones de transporte
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{"https://localhost:9200"},
		Username:  "admin", // For testing only. Don't store credentials in code.
		Password:  "Opensearch1234*",
	})
	if err != nil {
		log.Fatalf("Error creando el cliente de OpenSearch: %s", err)
	}

	//impresion de datos del cliente de opensearch
	clientres, err := client.Info()
	fmt.Printf("client: %v - error: %v", clientres, err)

	//pasamos el model a json para pasarlo como body de la request
	studentBytes, _ := json.Marshal(student)

	//creamos la request con indice students y el body que son los datos del estudiante
	req := opensearchapi.IndexRequest{
		Index: "students",
		Body:  bytes.NewReader(studentBytes),
	}
	//si tenemos un _id de documento lo agregamos a la request, si no opensearch nos crea uno
	//esto facilita la reutilizacion de codigo cuando el estudiante ya esta creado y cuando no
	if optionalID != "" {
		req.DocumentID = optionalID
	}

	res, _ := req.Do(context.Background(), client)
	defer res.Body.Close()
	return res.String(), nil
}

func SearchStudentByID(id int32) (mapper.StudentRequestModel, string, error) {
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{"https://localhost:9200"},
		Username:  "admin", // For testing only. Don't store credentials in code.
		Password:  "Opensearch1234*",
	})
	if err != nil {
		log.Fatalf("Error creando el cliente de OpenSearch: %s", err)
	}
	clientResponse, err := client.Info()
	fmt.Printf("client: %v - error: %v", clientResponse, err)

	fmt.Printf("id usada: %v \n", id)

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"id": map[string]interface{}{
					"value": id,
				},
			},
		},
	}

	content, _ := json.Marshal(query)

	req := opensearchapi.SearchRequest{
		Index: []string{"students"},
		Body:  bytes.NewReader(content),
	}

	res, _ := req.Do(context.Background(), client)

	var body OpenSearchResponse

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		log.Fatalf("error parseando resultado: %v", err)
	}

	defer res.Body.Close()
	return body.Hits.Hits[0].Source, body.Hits.Hits[0].ID, nil
}
