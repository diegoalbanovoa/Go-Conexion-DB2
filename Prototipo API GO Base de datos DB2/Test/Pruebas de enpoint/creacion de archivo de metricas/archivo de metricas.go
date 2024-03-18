package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Realizar una solicitud GET al endpoint /shutdown
	resp, err := http.Get("http://localhost:8080/shutdown")
	if err != nil {
		fmt.Println("Error al enviar la solicitud GET:", err)
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
}
