package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Datos de la consulta SQL
	query := "UPDATE CALIFIR SET CALSAR = 'M' WHERE TIPIDE = 45 AND NUMIDE = '000000000986' AND TIPCARD = 1;"

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
