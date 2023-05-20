package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MostrarTablaMultiplicar() {

	var numero int
	var err error
	scanner := bufio.NewScanner(os.Stdin)

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

	fmt.Printf("Tabla de Multiplicar del %d \n", numero)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", numero, i, (numero * i))
	}

}
