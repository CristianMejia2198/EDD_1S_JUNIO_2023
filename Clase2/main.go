package main

import (
	"Clase2/estructura"
)

func main() {
	listaSimple := &estructura.Lista{Inicio: nil, Longitud: 0}
	listaSimple.Insertar(11)
	listaSimple.Insertar(28)
	listaSimple.Insertar(13)
	listaSimple.Insertar(64)
	listaSimple.Insertar(95)
	listaSimple.Mostrar()
}
