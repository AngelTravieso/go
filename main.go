// Paquete principal
package main

import (
	// Paque que implementa funciones I/O
	"fmt"

	// Tiene la información del equipo donde se ejecuta el sistema
	"runtime"

	// Importar paquete 'variables'
	"github.com/AngelTravieso/go/variables"
)

// Punto de entrada del programa
func main() {
	// variables.MostrarEnteros()
	// variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(4)

	fmt.Println(estado)
	fmt.Println(texto)

	// os := runtime.GOOS

	// IF

	// asignacion; condicion
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es Windows, es", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	// Switch
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		// Formatear el texto
		fmt.Printf("%s \n", os)
	}

}
