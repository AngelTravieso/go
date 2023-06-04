// Paquete principal
package main

// Paque que implementa funciones I/O
// Tiene la información del equipo donde se ejecuta el sistema
// Importar paquete 'variables'
// "github.com/AngelTravieso/go/variables"
// Importar paquete 'ejercicios'
// "github.com/AngelTravieso/go/ejercicios"
// Importar paquete 'teclado'
import (

	// alias paquete
	// ejer "github.com/AngelTravieso/go/ejer_interfaces"
	// "github.com/AngelTravieso/go/modelos"
	// def "github.com/AngelTravieso/go/defer_panic"
	"fmt"

	routines "github.com/AngelTravieso/go/goroutines"
)

// Punto de entrada del programa
func main() {
	// variables.MostrarEnteros()
	// variables.RestoVariables()
	// estado, texto := variables.ConviertoaTexto(4)

	// fmt.Println(estado)
	// fmt.Println(texto)

	// os := runtime.GOOS

	// IF

	// asignacion; condicion
	// if os := runtime.GOOS; os == "linux" || os == "OS X." {
	// 	fmt.Println("Esto no es Windows, es", os)
	// } else {
	// 	fmt.Println("Esto es Windows")
	// }

	// Switch
	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	// Formatear el texto
	// 	fmt.Printf("%s \n", os)
	// }

	// numero, texto := ejercicios.StringtoInt("112")

	// fmt.Printf("%d %s", numero, texto)

	// teclado.IngresoNumeros()

	// iteraciones.Iterar()

	// fmt.Println(ejercicios.MostrarTablaMultiplicar())

	// files.GrabaTablaMultiplicar()
	// files.SumaTabla()
	// files.LeerArchivo()

	// funciones.Calculos()
	// funciones.LlamarClosure()
	// funciones.Exponencial(2)

	// arreglos_slices.MuestroArreglos()
	// arreglos_slices.MuestroSlice()
	// arreglos_slices.Capacidad()

	// mapas.MostrarMapas()

	// users.AltaUsuario()

	// Pedro := new(modelos.Hombre)
	// ejer.HumanosRespirando(Pedro)

	// Maria := new(modelos.Mujer)
	// ejer.HumanosRespirando(Maria)

	// def.VemosDefer()
	// def.EjemploPanic()

	// routines.MiNombreLento("Angel Travieso")

	// Ejecución asíncrona
	go routines.MiNombreLento("Angel Travieso")

	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)

}
