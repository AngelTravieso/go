package files

import (
	"fmt"
	"os"

	"github.com/AngelTravieso/go/ejercicios"
)

// Crea un archivo con la tabla de multiplicar
func GrabaTablaMultiplicar() {
	var texto string = ejercicios.MostrarTablaMultiplicar()

	// os.Create(ruta del archivo) crea un archivo
	// Si el archivo ya existe lo borra y lo vuelve a crear
	archivo, err := os.Create("./files/txt/tabla.txt")

	// Si hay un error
	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return
	}

	fmt.Println("Archivo creado exitosamente...")

	// Escribe en un archivo (ruta del archivo, contenido)
	fmt.Fprintln(archivo, texto)

	// Cerrar el archivo
	archivo.Close()
}

// func SumaTabla() {

// }
