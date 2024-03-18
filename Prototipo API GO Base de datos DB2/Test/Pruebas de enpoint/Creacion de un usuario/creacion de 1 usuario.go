package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	registro := map[string]interface{}{
		"tipide":  1,
		"numide":  "123456789",
		"tipcard": 1,
		"codseg":  1,
		"calsar":  "AB",
		"calapl":  "A",
		"calnum":  1,
		"calusr":  "Usuario",
		"calfua":  "20220310",
		"calhua":  "235959",
	}

	jsonBytes, err := json.Marshal(registro)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/createCalifir", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println("Error al enviar solicitud HTTP:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Respuesta del servidor:", resp.Status)
}
