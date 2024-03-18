package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
)

const (
	archivoEntrada = "C:\\Users\\diego\\Downloads\\Prototipo Api Python Base de Datos Oracle\\1. Creacion de usuarios en go\\usuarios_unicos.csv"
	archivoSalida  = "Nueva_calificacion_usuarios_de_prueba.csv"
)

func generarCalsar() string {
	opciones := []string{"A", "B", "C"}
	return opciones[rand.Intn(len(opciones))]
}

func eliminarDuplicados(registros [][]string) [][]string {
	// Crear un mapa para almacenar los registros únicos
	unicos := make(map[string]bool)
	result := [][]string{}

	// Iterar sobre los registros y agregarlos al mapa
	for _, registro := range registros {
		key := fmt.Sprintf("%s,%s,%s", registro[0], registro[1], registro[2])
		if !unicos[key] {
			unicos[key] = true
			result = append(result, registro)
		}
	}

	return result
}

func main() {
	// Solicitar semilla por consola
	var semilla int64
	fmt.Print("Ingresa la semilla para el generador de números aleatorios: ")
	_, err := fmt.Scanln(&semilla)
	if err != nil {
		fmt.Println("Semilla inválida.")
		return
	}
	rand.Seed(semilla)

	// Abrir archivo CSV de entrada
	archivoEntrada, err := os.Open(archivoEntrada)
	if err != nil {
		fmt.Println("Error al abrir el archivo de entrada:", err)
		return
	}
	defer archivoEntrada.Close()

	// Leer el archivo CSV y eliminar duplicados
	lector := csv.NewReader(archivoEntrada)
	registros, err := lector.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo CSV:", err)
		return
	}
	registros = eliminarDuplicados(registros)

	// Crear archivo CSV de salida
	archivoSalida, err := os.Create(archivoSalida)
	if err != nil {
		fmt.Println("Error al crear el archivo de salida:", err)
		return
	}
	defer archivoSalida.Close()

	// Crear escritor CSV
	escritor := csv.NewWriter(archivoSalida)
	defer escritor.Flush()

	// Escribir los registros en el archivo CSV de salida
	for _, registro := range registros {
		calsar := generarCalsar()
		registro = append(registro, calsar)
		if err := escritor.Write(registro); err != nil {
			fmt.Println("Error al escribir el registro en el archivo CSV de salida:", err)
			return
		}
	}

	fmt.Println("Archivo CSV de salida generado correctamente.")
}
