package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLento(nombre string, canal1 chan bool) {

	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		// Pausa de 1seg (lo maneja en milisegundos)
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}

	// asignar valor al channel
	canal1 <- true

}
