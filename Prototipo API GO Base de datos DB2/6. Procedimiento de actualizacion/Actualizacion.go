package main

import (
	"bytes"
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

		query := fmt.Sprintf("UPDATE CALIFIR SET CALSAR = '%s' WHERE TIPIDE = %d AND NUMIDE = '%s' AND TIPCARD = %d;", calsar, tipide, numide, tipcard)

		// Crear el cuerpo de la solicitud
		payload, err := json.Marshal(map[string]interface{}{
			"query": query,
		})
		if err != nil {
			fmt.Println("Error al convertir los datos a JSON:", err)
			return
		}

		// Realizar la solicitud POST al endpoint /update_calfir
		resp, err := http.Post("http://Localhost:8080/update_calfir", "application/json", bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("Error al enviar la solicitud POST:", err)
			return
		}
		defer resp.Body.Close()

		// Leer la respuesta del servidor
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error al leer la respuesta del servidor:", err)
			return
		}

		fmt.Println("Respuesta del servidor:", string(body))
	}
}
