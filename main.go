// Paquete principal
package main

import (
	// Paque que implementa funciones I/O
	"fmt"

	"github.com/AngelTravieso/go/variables"
)

// Punto de entrada del programa
func main() {
	// variables.MostrarEnteros()
	// variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(4)

	fmt.Println(estado)
	fmt.Println(texto)

}
