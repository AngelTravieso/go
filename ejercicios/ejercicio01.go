package ejercicios

import (
	"strconv"
)

// Parámetros:
//
// - text: cadena de texto númerica para convertir en int
//
// Valores de retorno:
//
// - int con el String equivalente a int
//
// - String con el mensaje descriptivo del número (es mayor a 100 o no)
func StringtoInt(text string) (int, string) {

	// _ para indicar que no necesito el otro dato que retorna
	// num, _ := strconv.Atoi(text)

	num, err := strconv.Atoi(text)

	if err != nil {
		// Error() devuelve mensaje de error (string)
		return 0, "Hubo un error " + err.Error()
	}

	if num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}

}
