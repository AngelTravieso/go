package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Devuelve la tabla de multiplicar de un número específico
// Valores de retorno:
//
// - string con la tabla de multiplicar
func TablaMultiplicar() string {

	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	var err error
	var texto string

	for {
		fmt.Println("Ingrese un numero por teclado: ")

		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())

			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	texto += fmt.Sprintf("Tabla de Multiplicar del %d \n", numero)

	for i := 1; i <= 10; i++ {
		// fmt.Printf("%d x %d = %d \n", numero, i, (numero * i))

		// Sprintf retorna el string mostrado por consola
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, (numero * i))

	}

	return texto

}
