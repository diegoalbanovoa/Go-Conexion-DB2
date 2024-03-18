// main_test.go
package main

import (
	"encoding/csv"
	"os"
	"strings"
	"testing"
)

func TestUsuariosUnicosEnCSV(t *testing.T) {
	archivoCSV, err := os.Open("usuarios_unicos.csv")
	if err != nil {
		t.Fatalf("Error al abrir el archivo CSV de prueba: %v", err)
	}
	defer archivoCSV.Close()

	lector := csv.NewReader(archivoCSV)
	usuariosLeidos, err := lector.ReadAll()
	if err != nil {
		t.Fatalf("Error al leer el archivo CSV de prueba: %v", err)
	}

	// Crear un mapa para almacenar los usuarios y verificar su unicidad
	usuariosMap := make(map[string]bool)
	for _, usuario := range usuariosLeidos {
		clave := strings.Join(usuario, ",")
		if usuariosMap[clave] {
			t.Errorf("Usuario duplicado encontrado: %s", clave)
		}
		usuariosMap[clave] = true
	}
}
