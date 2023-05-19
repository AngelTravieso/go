package variables

import "fmt"

// Para que la función sea pública se debe comenzar con una letra mayúscula

// La variable es exportada por defecto, ya que por tener la letra mayuscula es una funcion publica y es exportada
func MostrarEnteros() {

	// Las variables no se inicializan en null
	// Se inicializa con el valor más bajo permitido para el valor de la variable
	var intComunt int
	intde32 := int32(10)
	intde64 := int64(100)

	// Imprimir en consola
	fmt.Println("intcomun =", intComunt)
	fmt.Println("intde32 =", intde32)
	fmt.Println("intde64 =", intde64)

}
