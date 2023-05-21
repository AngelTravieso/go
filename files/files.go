package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AngelTravieso/go/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

// Crea un archivo con la tabla de multiplicar
func GrabaTablaMultiplicar() {
	var texto string = ejercicios.TablaMultiplicar()

	// os.Create(ruta del archivo) crea un archivo
	// Si el archivo ya existe lo borra y lo vuelve a crear
	archivo, err := os.Create(fileName)

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

func SumaTabla() {

	var texto string = ejercicios.TablaMultiplicar()
	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}

}

func Append(filen string, texto string) bool {

	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)

	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}

	arch.Close()
	return true

}

// Con ioutil
// func LeerArchivo() {
// 	archivo, err := ioutil.ReadFile(fileName)

// 	if err != nil {
// 		fmt.Println("Error al leer archivo " + err.Error())
// 		return
// 	}

// 	// archivo es un array de bytes, se hace un cast para imprimir el string
// 	fmt.Println(string(archivo))

// }

// Con os
func LeerArchivo() {
	archivo, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}

	// Los datos estan en el archivo
	scanner := bufio.NewScanner(archivo)

	// Mientras pueda escanear
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println(">" + registro)
	}

	archivo.Close()

}
