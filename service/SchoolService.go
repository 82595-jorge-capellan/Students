package SchoolService

import (
	"bytes"
	"io"
	"net/http"
	"fmt"
	"encoding/json"
	"log"

	pb "github.com/82595-jorge-capellan/protobuf"
)

// const (
// 	apiURL      = "https://api.jsonbin.io/v3/b/<BIN_ID>" // Reemplazar con tu BIN_ID
// 	apiKey      = "<YOUR_API_KEY>"                       // Reemplazar con tu API key
// 	contentType = "application/json"
// )

func AddStudent(jsonreq string) (*pb.StudentResponse, error) {
	url := "https://api.jsonbin.io/v3/b/682a35bf8a456b7966a092ff" // Reemplazá <BIN_ID>
	apiKey := "$2a$10$CpW4cEE3ebSGrj7erWp1OOq3mwiFsF8aB5yu/tl6fdkcDcGpujNAe" 

	originalJson, _ := GetJSON()
	printSliceOfMaps(originalJson)

	// 3. Agregar el objeto obtenido a la lista
	var objFromString map[string]interface{}
	err := json.Unmarshal([]byte(jsonreq), &objFromString)
	if err != nil {
    	panic(err)
	}
	originalJson = append(originalJson, objFromString)
	printSliceOfMaps(originalJson)

	jsonData, err := json.Marshal(originalJson)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creando request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error ejecutando request: %w", err)
	}
	defer resp.Body.Close()

	// Leer y mostrar respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo respuesta: %w", err)
	}

	fmt.Println("Response:", string(body))
	return &pb.StudentResponse{
		Status: string(body),
		FinalScore: 0,
		}, nil
}

func GetJSON() ([]map[string]interface{}, error) {
	url := "https://api.jsonbin.io/v3/b/682a35bf8a456b7966a092ff/latest"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creando request: %v", err)
	}

	// Agregar headers personalizados
	
	req.Header.Set("X-Master-Key", "$2a$10$CpW4cEE3ebSGrj7erWp1OOq3mwiFsF8aB5yu/tl6fdkcDcGpujNAe")
	req.Header.Set("X-Bin-Meta", "false")

	// Cliente HTTP estándar
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("estado no OK: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo body: %v", err)
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error parseando JSON: %v", err)
	}

	return result, nil
}

func printMap(m map[string]interface{}) {
    for key, value := range m {
        log.Printf("Clave: %s, Valor: %v\n", key, value)
    }
}

func printSliceOfMaps(slice []map[string]interface{}) {
    for i, m := range slice {
        log.Printf("Elemento %d:\n", i)
        for key, value := range m {
            log.Printf("  %s: %v\n", key, value)
        }
    }
}



// func AddStudent(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPut {
// 		http.Error(w, "Sólo se acepta PUT", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Lee el body recibido
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Error leyendo body", http.StatusBadRequest)
// 		return
// 	}
// 	defer r.Body.Close()

// 	// Prepara el request hacia jsonbin
// 	req, err := http.NewRequest(http.MethodPut, apiURL, bytes.NewBuffer(body))
// 	if err != nil {
// 		http.Error(w, "Error creando solicitud externa", http.StatusInternalServerError)
// 		return
// 	}
// 	req.Header.Set("Content-Type", contentType)
// 	req.Header.Set("X-Master-key", apiKey)

// 	// Ejecuta el request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		http.Error(w, "Error al hacer PUT externo", http.StatusInternalServerError)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Retorna lo que respondió jsonbin
// 	respBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		http.Error(w, "Error leyendo respuesta externa", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", contentType)
// 	w.WriteHeader(resp.StatusCode)
// 	w.Write(respBody)
// }

// func main() {
// 	http.HandleFunc("/proxy", handler)

// 	log.Println("Escuchando en http://localhost:8080/proxy")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }