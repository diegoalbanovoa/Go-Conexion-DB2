package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/califirnew")
	if err != nil {
		fmt.Println("Error al obtener datos de CALIFIRNEW:", err)
		return
	}
	defer resp.Body.Close()

	// Leer la respuesta del servidor
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta del servidor:", err)
		return
	}

	// Mostrar la respuesta del servidor
	fmt.Println("Respuesta del servidor:", string(body))

	// Decodificar la respuesta JSON en una estructura de datos
	var califirnewData []map[string]interface{}
	if err := json.Unmarshal(body, &califirnewData); err != nil {
		fmt.Println("Error al decodificar respuesta JSON de CALIFIRNEW:", err)
		return
	}

	// Mostrar la longitud de califirnewData
	fmt.Println("Longitud de califirnewData:", len(califirnewData))

	// Iterar sobre los registros de CALIFIRNEW y mostrarlos por consola
	for _, data := range califirnewData {
		tipide := int(data["tipide"].(float64)) // Convertir a int
		numide := data["numide"].(string)
		tipcard := int(data["tipcard"].(float64)) // Convertir a int
		calsar := data["calsar"].(string)

		fmt.Printf("Obteniendo registro: TIPIDE=%d, NUMIDE=%s, TIPCARD=%d\n", tipide, numide, tipcard, calsar)
	}
}
