package main

import (
	"Clase3/estructuras"
	"fmt"
)

func main() {
	Listadoble := &estructuras.ListaDoble{Inicio: nil, Longitud: 0}
	Listadoble.Insertar(2017000000, "Cristian")
	Listadoble.Insertar(2017000001, "Alberto")
	Listadoble.Insertar(2017000002, "Suy")
	Listadoble.Insertar(2017000003, "Mejia")
	fmt.Println("Ascendente")
	Listadoble.MostrarAscendente()
	fmt.Println("Descendente")
	Listadoble.MostrarDescente()
	Listadoble.Reporte()
}
