package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {

	// defer indica al sistema cual es la instruccion final de la ejecución, puedo tener n defers en el código

	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")

	/*
		Cuando tenemos multiples defer en un programa, el orden de ejecucion de las instrucciones defer será
		LIFO (Last In First Out), del ultimo al primero

		defer fmt.Println("Uno")
		defer fmt.Println("Dos")
		defer fmt.Println("Tres")


	*/

	fmt.Println("Este es el segundo mensaje")

}

func EjemploPanic() {

	// el defer de esta manera devuelve una funcion
	defer func() {
		// recover se usa con defer (si o si)
		// recover previene el final del programa y recupera el programa de un "panic"
		reco := recover()

		// Hubo un error
		if reco != nil {
			// Fatalf es un equivalente a PrintF y termina el programa so.Exit(1)
			log.Fatalf("Ocurrió un error que generó Panic \n %s", reco)
		}
	}()

	// panic aborta el programa inmediatamente con un mensaje que envia a consola, es cuando hay una falla
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}

}
