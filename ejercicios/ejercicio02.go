package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MostrarTablaMultiplicar() string {

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
