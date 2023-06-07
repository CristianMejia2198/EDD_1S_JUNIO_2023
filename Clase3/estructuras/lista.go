package estructuras

import (
	"fmt"
	"strconv"
)

type ListaDoble struct {
	Inicio   *NodoLista
	Longitud int
}

func (l *ListaDoble) estaVacia() bool {
	return l.Longitud == 0
}

func (l *ListaDoble) Insertar(carnet int, nombre string) {
	nuevoAlumno := &Alumno{Carnet: carnet, Nombre: nombre}
	// 0x8494AB8 = {20200000, Jose}
	//nuevoAlumno := &Alumno{carnet, nombre}
	if l.estaVacia() {
		l.Inicio = &NodoLista{nuevoAlumno, nil, nil}
		l.Longitud++
	} else {
		aux := l.Inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = &NodoLista{nuevoAlumno, nil, aux}
		l.Longitud++
	}
}

func (l *ListaDoble) MostrarAscendente() {
	aux := l.Inicio
	for aux != nil {
		fmt.Print(aux.alumno.Carnet)
		fmt.Println(" --> ", aux.alumno.Nombre)
		aux = aux.siguiente
	}
}

func (l *ListaDoble) MostrarDescente() {
	aux := l.Inicio
	for aux.siguiente != nil {
		aux = aux.siguiente
	}
	/*Imprimir hacia atras*/
	// Inicio = {Alumno, 0x0000001, nil}
	for aux != nil {
		fmt.Print(aux.alumno.Carnet)
		fmt.Println(" --> ", aux.alumno.Nombre)
		aux = aux.anterior
	}
}

func (l *ListaDoble) Reporte() {
	nombreArchivo := "./listadoble.dot"
	nombreImagen := "./listadoble.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := l.Inicio
	contador := 0
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < l.Longitud; i++ {
		texto += "nodo" + strconv.Itoa(i) + "[label=\"" + aux.alumno.Nombre + "\"];\n"
		aux = aux.siguiente
	}
	for i := 0; i < l.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	crearArchivo(nombreArchivo)
	escribirArchivo(texto, nombreArchivo)
	ejecutar(nombreImagen, nombreArchivo)
}
