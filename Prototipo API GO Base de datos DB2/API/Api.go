package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/ibmdb/go_ibm_db"
)

type Metrica struct {
	URL             string        `json:"url"`
	Metodo          string        `json:"metodo"`
	TiempoRespuesta time.Duration `json:"tiempo_respuesta"`
}

var metricas []Metrica

func metricasMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		tiempoRespuesta := time.Since(start)
		metrica := Metrica{
			URL:             c.Request.URL.Path,
			Metodo:          c.Request.Method,
			TiempoRespuesta: tiempoRespuesta,
		}
		metricas = append(metricas, metrica)
	}
}

type Califir struct {
	Tipide  int    `json:"tipide"`
	Numide  string `json:"numide"`
	Tipcard int    `json:"tipcard"`
	Codseg  int    `json:"codseg"`
	Calsar  string `json:"calsar"`
	Calapl  string `json:"calapl"`
	Calnum  int    `json:"calnum"`
	Calusr  string `json:"calusr"`
	Calfua  string `json:"calfua"`
	Calhua  string `json:"calhua"`
}

var califirData []Califir

type Califirnew struct {
	Tipide  int     `json:"tipide"`
	Numide  string  `json:"numide"`
	Tipcard int     `json:"tipcard"`
	Calsar  *string `json:"calsar,omitempty"`
}

var califirnewData []Califirnew

func main() {
	r := gin.Default()

	r.Use(metricasMiddleware())

	r.POST("/createCalifir", createCalifirHandler)
	r.POST("/createCalifirnew", createCalifirnewHandler)
	r.GET("/shutdown", shutdownHandler)
	r.GET("/califirnew", califirnewHandler)
	r.POST("/update_calfir", updateCalfirHandler)

	r.Run(":8080")
}

func createCalifirHandler(c *gin.Context) {
	start := time.Now()
	var registro Califir
	if err := c.BindJSON(&registro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	duration := time.Since(start)
	metrica := Metrica{
		URL:             "/createCalifir",
		Metodo:          "POST",
		TiempoRespuesta: duration,
	}
	metricas = append(metricas, metrica)
	connString := "DATABASE=testdb;HOSTNAME=localhost;PORT=50000;PROTOCOL=TCPIP;UID=db2inst1;PWD=password;"
	db, err := sql.Open("go_ibm_db", connString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error conectando a la base de datos"})
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hacer ping a la base de datos"})
		return
	}
	_, err = db.Exec("INSERT INTO CALIFIR (TIPIDE, NUMIDE, TIPCARD, CODSEG, CALSAR, CALAPL, CALNUM, CALUSR, CALFUA, CALHUA) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		registro.Tipide, registro.Numide, registro.Tipcard, registro.Codseg, registro.Calsar, registro.Calapl, registro.Calnum, registro.Calusr, registro.Calfua, registro.Calhua)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar registro"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registro insertado correctamente"})
}

func createCalifirnewHandler(c *gin.Context) {
	start := time.Now()
	var registro Califirnew
	if err := c.BindJSON(&registro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	duration := time.Since(start)
	metrica := Metrica{
		URL:             "/createCalifir",
		Metodo:          "POST",
		TiempoRespuesta: duration,
	}
	metricas = append(metricas, metrica)
	connString := "DATABASE=testdb;HOSTNAME=localhost;PORT=50000;PROTOCOL=TCPIP;UID=db2inst1;PWD=password;"
	db, err := sql.Open("go_ibm_db", connString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error conectando a la base de datos"})
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hacer ping a la base de datos"})
		return
	}
	_, err = db.Exec("INSERT INTO CALIFIRNEW (TIPIDE, NUMIDE, TIPCARD, CALSAR) VALUES (?, ?, ?, ?)",
		registro.Tipide, registro.Numide, registro.Tipcard, registro.Calsar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar registro"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registro insertado correctamente"})
}

func shutdownHandler(c *gin.Context) {
	guardarMetricasEnArchivo()
	c.JSON(http.StatusOK, gin.H{"message": "API apagada. Métricas guardadas en el archivo metricas.txt"})
}

func califirnewHandler(c *gin.Context) {
	// Conexión a la base de datos
	connString := "DATABASE=testdb;HOSTNAME=localhost;PORT=50000;PROTOCOL=TCPIP;UID=db2inst1;PWD=password;"
	db, err := sql.Open("go_ibm_db", connString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error conectando a la base de datos"})
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hacer ping a la base de datos"})
		return
	}

	// Consultar todos los registros de CALIFIRNEW
	rows, err := db.Query("SELECT TIPIDE, NUMIDE, TIPCARD, CALSAR FROM CALIFIRNEW")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener registros de CALIFIRNEW"})
		return
	}
	defer rows.Close()

	// Crear una estructura para almacenar los datos de CALIFIRNEW
	var califirnewData []Califirnew
	for rows.Next() {
		var data Califirnew
		if err := rows.Scan(&data.Tipide, &data.Numide, &data.Tipcard, &data.Calsar); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escanear registros de CALIFIRNEW"})
			return
		}
		califirnewData = append(califirnewData, data)
	}

	c.JSON(http.StatusOK, califirnewData)
}

func updateCalfirHandler(c *gin.Context) {
	var req struct {
		Query string `json:"query"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Conexión a la base de datos
	connString := "DATABASE=testdb;HOSTNAME=localhost;PORT=50000;PROTOCOL=TCPIP;UID=db2inst1;PWD=password;"
	db, err := sql.Open("go_ibm_db", connString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error conectando a la base de datos"})
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hacer ping a la base de datos"})
		return
	}

	// Ejecutar la consulta SQL recibida
	_, err = db.Exec(req.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al ejecutar la consulta SQL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Consulta ejecutada correctamente"})
}

func guardarMetricasEnArchivo() {
	file, err := os.Create("metricas.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo de métricas:", err)
		return
	}
	defer file.Close()

	for _, metrica := range metricas {
		line := fmt.Sprintf("URL: %s, Método: %s, Tiempo de respuesta: %s\n", metrica.URL, metrica.Metodo, metrica.TiempoRespuesta.String())
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Error al escribir en el archivo de métricas:", err)
			return
		}
	}
	fmt.Println("Métricas guardadas en el archivo metricas.txt")
}

func shutdown() {
	guardarMetricasEnArchivo()
	os.Exit(0)
}
