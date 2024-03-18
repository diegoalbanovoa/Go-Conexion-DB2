package main

import (
	"encoding/csv"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"math/rand"
	"os"
	"strconv"
)

const (
	nombreArchivo = "usuarios_unicos.csv"
)

func generarUsuariosUnicos() [][]string {
	usuariosUnicos := [][]string{}
	for i := 1; i <= 99; i++ {
		tipide := strconv.Itoa(i)
		for j := 0; j < 1000; j++ {
			numide := fmt.Sprintf("%012d", j+1)
			for _, tipcard := range []string{"1", "2", "3"} {
				usuariosUnicos = append(usuariosUnicos, []string{tipide, numide, tipcard})
			}
		}
	}
	return usuariosUnicos
}

func main() {
	var cantidadUsuarios int
	fmt.Print("¿Cuántos usuarios deseas generar?: ")
	_, err := fmt.Scanln(&cantidadUsuarios)
	if err != nil || cantidadUsuarios <= 0 {
		fmt.Println("Cantidad de usuarios inválida.")
		return
	}

	var semilla int64
	fmt.Print("Ingresa la semilla para la generación de números aleatorios: ")
	_, err = fmt.Scanln(&semilla)
	if err != nil {
		fmt.Println("Semilla inválida.")
		return
	}

	r := rand.New(rand.NewSource(semilla))
	usuariosUnicos := generarUsuariosUnicos()
	usuarios := make([][]string, 0)
	bar := progressbar.Default(int64(cantidadUsuarios))
	for i := 0; i < cantidadUsuarios; i++ {
		idx := r.Intn(len(usuariosUnicos))
		usuario := usuariosUnicos[idx]
		usuarios = append(usuarios, usuario)
		usuariosUnicos = append(usuariosUnicos[:idx], usuariosUnicos[idx+1:]...)
		bar.Add(1)
	}

	archivoCSV, err := os.Create(nombreArchivo)
	if err != nil {
		fmt.Println("Error al crear el archivo CSV:", err)
		return
	}
	defer archivoCSV.Close()

	escritor := csv.NewWriter(archivoCSV)
	defer escritor.Flush()

	for _, usuario := range usuarios {
		if err := escritor.Write(usuario); err != nil {
			fmt.Println("Error al escribir en el archivo CSV:", err)
			return
		}
	}

	fmt.Printf("\nArchivo CSV generado: %s (%d usuarios)\n", nombreArchivo, cantidadUsuarios)
}
