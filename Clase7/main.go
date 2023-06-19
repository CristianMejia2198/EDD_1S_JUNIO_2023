package main

import "Clase7/estructuras"

func main() {
	arbol := &estructuras.Arbol{Raiz: nil}
	arbol.InsertarElemento(10)
	arbol.InsertarElemento(8)
	arbol.InsertarElemento(11)
	arbol.InsertarElemento(20)
	arbol.InsertarElemento(1)
	arbol.InsertarElemento(4)
	arbol.InsertarElemento(9)
	arbol.InsertarElemento(7)
	arbol.GenerarRecorrido()

	/*
					10
				8		11
			7		9		20
		1
			4
	*/
}
