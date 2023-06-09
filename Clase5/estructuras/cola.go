package estructuras

import (
	"fmt"
	"strconv"
)

type Cola struct {
	Primero  *NodoCola
	Longitud int
}

func (c *Cola) Encolar(carnet string, nombre string) {
	nuevoAlumno := &Alumno{Carnet: carnet, Nombre: nombre}
	if c.Longitud == 0 {
		nuevoNodo := &NodoCola{nuevoAlumno, nil}
		c.Primero = nuevoNodo
		c.Longitud++
	} else {
		nuevoNodo := &NodoCola{nuevoAlumno, nil}
		aux := c.Primero
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = nuevoNodo
		c.Longitud++
	}
}

func (c *Cola) Descolar() {
	if c.Longitud == 0 {
		fmt.Println("No hay alumnos pendientes en la cola")
	} else {
		c.Primero = c.Primero.Siguiente
		c.Longitud--
	}
}

func (c *Cola) Graficar() {
	nombre_archivo := "./cola.dot"
	nombre_imagen := "cola.jpg"
	texto := "digraph cola{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull2[label=\"null\"];\n"
	aux := c.Primero
	contador := 0
	for i := 0; i < c.Longitud; i++ {
		texto = texto + "nodo" + strconv.Itoa(i) + "[label=\"{Nombre: " + aux.Alumno.Nombre + ", Carnet: " + aux.Alumno.Carnet + "|}\"];\n"
		aux = aux.Siguiente
	}
	for i := 0; i < c.Longitud-1; i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		contador = c
	}
	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"
	crearArchivo(nombre_archivo)
	escribirArchivo(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}
