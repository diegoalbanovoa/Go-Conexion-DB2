package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Abrir el archivo CSV
	file, err := os.Open("C:\\Users\\diego\\Downloads\\Prototipo API GO Base de datos DB2\\2. Creacion de registros en go\\base_de_registros_de_prueba.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Crear un lector CSV
	reader := csv.NewReader(file)

	// Leer todas las líneas del archivo
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo CSV:", err)
		return
	}

	// Iterar sobre las líneas del archivo
	for _, line := range lines {
		// Convertir los datos a los tipos de datos correctos
		tipide, _ := strconv.Atoi(line[0])
		numide := line[1]
		tipcard, _ := strconv.Atoi(line[2])
		codseg, _ := strconv.Atoi(line[3])
		calsar := line[4]
		calapl := line[5]
		calnum, _ := strconv.Atoi(line[6])
		calusr := line[7]
		calfua := line[8]
		calhua := line[9]

		// Crear un mapa para el registro
		registro := map[string]interface{}{
			"tipide":  tipide,
			"numide":  numide,
			"tipcard": tipcard,
			"codseg":  codseg,
			"calsar":  calsar,
			"calapl":  calapl,
			"calnum":  calnum,
			"calusr":  calusr,
			"calfua":  calfua,
			"calhua":  calhua,
		}

		// Convertir el registro a JSON
		jsonBytes, err := json.Marshal(registro)
		if err != nil {
			fmt.Println("Error al convertir a JSON:", err)
			return
		}

		// Enviar el registro al servidor
		resp, err := http.Post("http://localhost:8080/createCalifir", "application/json", bytes.NewBuffer(jsonBytes))
		if err != nil {
			fmt.Println("Error al enviar solicitud HTTP:", err)
			continue
		}
		defer resp.Body.Close()

		fmt.Println("Respuesta del servidor:", resp.Status)
	}
}
