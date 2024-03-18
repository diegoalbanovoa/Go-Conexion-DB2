package main

import (
	_ "bytes"
	"encoding/csv"
	_ "encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	archivoCSV = "C:\\Users\\diego\\Downloads\\Prototipo Api Python Base de Datos Oracle\\4. Creacion de datos de prueba a actualizar\\Nueva_calificacion_usuarios_de_prueba.csv"
	baseURL    = "http://localhost:8080/"
)

func main() {
	// Abrir archivo CSV
	file, err := os.Open(archivoCSV)
	if err != nil {
		fmt.Println("Error al abrir el archivo CSV:", err)
		return
	}
	defer file.Close()

	// Crear cliente HTTP
	client := &http.Client{}

	// Leer el archivo CSV y enviar cada fila como una solicitud POST al API
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// Convertir los datos a los tipos de datos correctos
		tipide, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println("Error al convertir tipide a entero:", err)
			continue
		}
		numide := record[1]
		tipcard, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Println("Error al convertir tipcard a entero:", err)
			continue
		}
		calsar := record[3]

		// Crear datos para la solicitud POST
		data := strings.NewReader(fmt.Sprintf(`{
			"tipide": %d,
			"numide": "%s",
			"tipcard": %d,
			"calsar": "%s"
		}`, tipide, numide, tipcard, calsar))

		// Crear solicitud POST al API
		req, err := http.NewRequest("POST", baseURL+"createCalifirnew", data)
		if err != nil {
			fmt.Println("Error al crear la solicitud HTTP:", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		// Enviar solicitud al API
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error al enviar la solicitud HTTP:", err)
			return
		}
		defer resp.Body.Close()

		// Verificar respuesta del API
		if resp.StatusCode == http.StatusOK {
			fmt.Println("Registro creado correctamente")
		} else {
			fmt.Println("Error al crear el registro:", resp.StatusCode)
		}
	}
}
