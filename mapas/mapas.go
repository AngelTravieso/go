package mapas

import "fmt"

func MostrarMapas() {

	// map[clave_tipo]valor_tipo, largo determinado
	paises := make(map[string]string)

	// Devuelve un mapa vacio
	fmt.Println(paises)

	paises["Venezuela"] = "Caracas"
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)

	// Acceder a un valor
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       20,
		"Boca Juniors": 16,
	}

	fmt.Println(campeonato)

	// Recorrer map, el range es como un forEach
	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	// Borrar elemento de un mapa
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Juventus"]

	// %t imprimir bool en printF
	fmt.Printf("El puntaje capturado es %d y el equipo existe = %t", puntaje, existe)

	// Si el map no existe devuelve el valor mínimo
	// campeonato["Juventus"] = 0, false

}
