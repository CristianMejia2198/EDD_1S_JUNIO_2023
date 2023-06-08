package main

import (
	"Clase4/estructuras"
	"fmt"
)

func main() {
	listaCircular := &estructuras.ListaCircular{Inicio: nil, Longitud: 0}
	estructuras.LeerArchivo("Estudiante.csv", listaCircular)
	listaCircular.Mostrar()

	/*Extra*/
	opcion := 0
	salir := false

	for !salir {
		fmt.Scanln(&opcion)
		fmt.Println("El numero que coloque fue ", opcion)
		if opcion == 2 {
			salir = true
		}
	}
}
