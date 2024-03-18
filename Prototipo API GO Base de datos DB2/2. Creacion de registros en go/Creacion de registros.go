package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
)

const (
	nombreArchivoUsuarios  = "C:\\Users\\diego\\Downloads\\Prototipo API GO Base de datos DB2\\1. Creacion de usuarios en go\\usuarios_unicos.csv"
	nombreArchivoRegistros = "base_de_registros_de_prueba.csv"
	mensajeSemilla         = "Ingresa la semilla para el generador de números aleatorios: "
)

func generarRegistros(usuario []string, semilla int64) []string {
	rand.Seed(semilla)
	registros := make([]string, 0, 8)
	for i := 0; i < 8; i++ {
		codseg := strconv.Itoa(rand.Intn(899) + 100)
		calsar := string([]rune{'A', 'B', 'C'}[rand.Intn(3)])
		calapl := string([]rune{'0', '1'}[rand.Intn(2)])
		calnum := strconv.Itoa(rand.Intn(90) + 10)
		calusr := "USR" + strconv.Itoa(rand.Intn(9000)+1000)
		fechaAleatoria := time.Now().UTC()
		calhua := fechaAleatoria.Format("150405")
		calfua := fechaAleatoria.Format("20060102")
		registro := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s", usuario[0], usuario[1], usuario[2], codseg, calsar, calapl, calnum, calusr, calfua, calhua)
		registros = append(registros, registro)
	}
	return registros
}

func main() {
	// Solicitar semilla por consola
	var semilla int64
	fmt.Print(mensajeSemilla)
	_, err := fmt.Scanln(&semilla)
	if err != nil {
		fmt.Println("Semilla inválida.")
		return
	}

	// Leer archivo CSV de usuarios
	archivoUsuarios, err := os.Open(nombreArchivoUsuarios)
	if err != nil {
		fmt.Println("Error al abrir el archivo de usuarios:", err)
		return
	}
	defer archivoUsuarios.Close()

	lector := csv.NewReader(archivoUsuarios)
	_, _ = lector.Read() // Omitir encabezado
	usuarios, err := lector.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo CSV de usuarios:", err)
		return
	}

	// Generar registros para cada usuario
	registros := make([]string, 0, len(usuarios)*8)
	bar := progressbar.Default(int64(len(usuarios) * 8))
	for _, usuario := range usuarios {
		registros = append(registros, generarRegistros(usuario, semilla)...)
		bar.Add(8)
	}
	bar.Finish()

	// Escribir registros en archivo CSV
	archivoRegistros, err := os.Create(nombreArchivoRegistros)
	if err != nil {
		fmt.Println("Error al crear el archivo CSV de registros:", err)
		return
	}
	defer archivoRegistros.Close()

	escritor := csv.NewWriter(archivoRegistros)
	defer escritor.Flush()

	for _, registro := range registros {
		if err := escritor.Write(strings.Split(registro, ",")); err != nil {
			fmt.Println("Error al escribir en el archivo CSV de registros:", err)
			return
		}
	}

	fmt.Printf("Archivo CSV de registros generado: %s\n", nombreArchivoRegistros)
}
