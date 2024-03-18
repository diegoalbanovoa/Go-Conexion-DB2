package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// Datos del registro a insertar
	registro := map[string]interface{}{
		"tipide":  1,
		"numide":  "123456789",
		"tipcard": 1,
		"calsar":  "AB",
	}

	// Convertir el registro a JSON
	jsonBytes, err := json.Marshal(registro)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	// Enviar el registro al servidor
	resp, err := http.Post("http://localhost:8080/createCalifirnew", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println("Error al enviar solicitud HTTP:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Respuesta del servidor:", resp.Status)
}
