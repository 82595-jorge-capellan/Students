package SchoolService

import (
	"bytes"
	"io"
	"net/http"
	"fmt"
	"encoding/json"
	"log"
)

func AddStudent(jsonreq []byte) ([]byte, error) {

	//secretos hardcodeados D:
	url := "https://api.jsonbin.io/v3/b/682a35bf8a456b7966a092ff"
	apiKey := "$2a$10$CpW4cEE3ebSGrj7erWp1OOq3mwiFsF8aB5yu/tl6fdkcDcGpujNAe" 

	// creamos la request de PUT con la url de la Bin y el json a enviar como array de bytes
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonreq))
	if err != nil {
		return nil, fmt.Errorf("error creando request: %w", err)
	}

	//seteamos headers como indica la api de jsonBin
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", apiKey)

	//creamos el cliente http y le hacemos ejecutar la request
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
	return body, nil
}

func GetJSON() ([]map[string]interface{}, error) {

	//mas secretos hardcodeado :(
	url := "https://api.jsonbin.io/v3/b/682a35bf8a456b7966a092ff/latest"

	//creamos la request con el metodo GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creando request: %v", err)
	}

	//seteamos headers para no traer metadata que ensucie el json final, (no se como quitarla)(probablemente con marshalling y la struct)
	req.Header.Set("X-Master-Key", "$2a$10$CpW4cEE3ebSGrj7erWp1OOq3mwiFsF8aB5yu/tl6fdkcDcGpujNAe")
	req.Header.Set("X-Bin-Meta", "false")

	// creamos el cliente y le hacemos ejecutar la request
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

func PrintMap(m map[string]interface{}) {
    for key, value := range m {
        log.Printf("Clave: %s, Valor: %v\n", key, value)
    }
}

func PrintSliceOfMaps(slice []map[string]interface{}) {
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

// originalJson, _ := GetJSON()
	// log.Println("data repetida(?)")
	// PrintSliceOfMaps(originalJson)

	// // 3. Agregar el objeto obtenido a la lista
	// var objFromString map[string]interface{}
	// err := json.Unmarshal([]byte(jsonreq), &objFromString)
	// if err != nil {
    // 	panic(err)
	// }
	// originalJson = append(originalJson, objFromString)
	// PrintSliceOfMaps(originalJson)

	// jsonData, err := json.Marshal(originalJson)
	// if err != nil {
	// 	panic(err)
	// }