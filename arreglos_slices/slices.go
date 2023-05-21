package arreglos_slices

import "fmt"

var tablas []int = []int{2, 5, 4}

var arreglo [10]int = [10]int{
	6,
	98,
	2,
	4,
	61,
	23,
	54312,
	8658,
	324,
}

func MuestroSlice() {

	fmt.Println(tablas)

	// Del elemento 3 al final
	porcion := arreglo[3:]

	// Desde el comienzo hasta la posicion 5
	porcion2 := arreglo[:5]

	// Del item 6 al 7
	porcion3 := arreglo[6:7]

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

func Capacidad() {

	// Tipo, elementos, capacidad
	elementos := make([]int, 5, 20)

	// cap() para obtener la capacidad declarada

	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)

	// Agregar elemento a un slice
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))

}
