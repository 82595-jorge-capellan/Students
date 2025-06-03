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

type MSearchResponse struct {
	Responses []struct {
		Hits struct {
			Hits []struct {
				Source mapper.StudentRequestModel `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	} `json:"responses"`
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
		Index: student.Subject,
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

func SearchStudentByID(id int32, subject string) (mapper.StudentRequestModel, string, error) {
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
	// clientResponse, err := client.Info()
	// fmt.Printf("client: %v - error: %v", clientResponse, err)

	// fmt.Printf("id usada: %v \n", id)

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
		Index: []string{subject},
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

func MsearchSearchStudent(id int32) []mapper.StudentRequestModel {
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
	// clientResponse, err := client.Info()
	// fmt.Printf("client: %v - error: %v", clientResponse, err)

	fmt.Printf("id usada: %v \n", id)

	indices := []string{"math", "biology", "chemistry"}
	queryMap := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"id": map[string]interface{}{
					"value": id,
				},
			},
		},
	}
	var buffer bytes.Buffer
	for _, index := range indices {
		// Línea 1: metadatos
		meta := map[string]interface{}{
			"index": index,
		}
		metaJSON, _ := json.Marshal(meta)
		buffer.Write(metaJSON)
		buffer.WriteByte('\n')

		// Línea 2: query
		queryJSON, _ := json.Marshal(queryMap)
		buffer.Write(queryJSON)
		buffer.WriteByte('\n')
	}

	req := opensearchapi.MsearchRequest{
		Body: bytes.NewReader(buffer.Bytes()),
	}

	res, _ := req.Do(context.Background(), client)

	var msearchResp MSearchResponse

	if err := json.NewDecoder(res.Body).Decode(&msearchResp); err != nil {
		log.Fatalf("Error decoding MSearch response: %s", err)
	}
	var results []mapper.StudentRequestModel

	// Extraer todos los _source de cada búsqueda
	for _, resp := range msearchResp.Responses {
		for _, hit := range resp.Hits.Hits {
			results = append(results, hit.Source)
		}
	}

	defer res.Body.Close()
	return results
}
