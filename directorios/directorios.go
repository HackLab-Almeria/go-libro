package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Comprobando directorios introducidos por parámetro")
	//Recibimos un array de parámetros, la primera posición es la ruta del script
	if len(os.Args) < 2 {
		println("Debes introducir los directorios")
		os.Exit(2)
	}
	directorios := os.Args
	// Valdiamos los directorios
	for _, dir := range directorios{
		if _, err := os.Stat(dir); os.IsNotExist(err){
			println("El directorio " + dir + " no existe")
		} else {
			println("El directorio" + dir + " existe")
		}

	}
	//https://stackoverflow.com/questions/18973335/golang-check-number-of-arguments-also-user-input-check-for-return-key-blan
	}
