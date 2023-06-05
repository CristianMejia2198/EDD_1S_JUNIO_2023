package main

import (
	"Clase1/variables"
	"fmt"
)

type alumno struct {
	carnet int
	nombre string
}

func main() {
	var x int = 1
	var y int = 2
	var mi [4]int
	variables.Sumar(x, y)
	u := 9
	mi_lista := [4]int{1, 2, 3, 4}
	fmt.Println(x)
	fmt.Println(u)
	mi[1] = 5
	fmt.Println(mi_lista)
	fmt.Println(mi)

	/*Nueva seccion*/
	var persona *alumno = &alumno{carnet: 2017, nombre: "Cristian"}
	fmt.Println("--------------------------------------------")
	fmt.Println("Por Valor")
	mostrar(y)
	fmt.Println(y)
	fmt.Println("Por Referencia")
	cambiar(persona)
	fmt.Println(*persona)
}

func mostrar(x int) {
	(x)++
	fmt.Println(x)
}

func cambiar(alum *alumno) {
	alum.carnet = 2020
}
