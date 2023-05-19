package variables

import (
	// Funciones I/O
	"fmt"
	// Para manejar fechas
	"time"
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
