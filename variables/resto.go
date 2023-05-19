package variables

import (
	// Funciones I/O
	"fmt"
	// Para manejar fechas
	"time"
	// Para convertir int a string
	"strconv"
)

var Nombre string
var Estado bool
var Sueldo float32 // float64
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()

	// El argumento de Println es <any>
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoaTexto(numero int) (bool, string) {
	// Itoa: Integer to alphanumeric
	texto := strconv.Itoa(numero)
	return true, texto
}
