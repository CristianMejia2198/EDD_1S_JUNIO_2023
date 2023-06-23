package estructuras

import (
	"strconv"
)

type Grafo struct {
	Principal *NodoMatrizDeAdyacencia
}

/*
Empleado 1234 -> 8995 -> 8596 -> 7983
8995 -> imagen1 -> filtro1 -> imagen8 -> filtro4
8596 -> imagen2 -> filtro1 - filtro2
7983 -> imagen3 -> filtro 4
*/

func (g *Grafo) insertarC(padre string, hijo string, filtro string) { // cliente, imagen
	nuevoNodo := &NodoMatrizDeAdyacencia{Valor: hijo}
	if g.Principal != nil && padre == g.Principal.Valor {
		aux := g.Principal
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = nuevoNodo
	} else {
		g.insertarF(padre)
		aux := g.Principal
		for aux != nil {
			if aux.Valor == padre {
				break
			}
			aux = aux.Abajo
		}
		if aux != nil {
			nuevoNodo.Siguiente = &NodoMatrizDeAdyacencia{Valor: filtro}
			aux.Siguiente = nuevoNodo
		}
	}
}

func (g *Grafo) insertarF(padre string) {
	nuevoNodo := &NodoMatrizDeAdyacencia{Valor: padre}
	if g.Principal == nil {
		g.Principal = nuevoNodo
	} else {
		aux := g.Principal
		for aux.Abajo != nil {
			if aux.Valor == padre {
				return
			}
			aux = aux.Abajo
		}
		aux.Abajo = nuevoNodo
	}
}

/*
{
    "Padre": "Empleado 1234",
    "Cliente": "1473",
    "Imagen": "tux",
    "Filtros": "Negativo"

	Principal: "Empleado 1234"
	Principal.Siguiente: "1473"
	Principal.Abajo: "1473"
	Principal.Abajo.Siguiente: "tux" (aux)
	aux.siguiente : "Negativo"
}
*/

func (g *Grafo) InsertarValores(padre string, cliente string, imagen string, filtros string) {
	if g.Principal == nil {
		g.insertarF(padre)
		g.insertarC(padre, cliente, "")
		g.insertarC(cliente, imagen, filtros)
	} else {
		g.insertarC(padre, cliente, "")
		g.insertarC(cliente, imagen, filtros)
	}
}

func (g *Grafo) Reporte() {
	cadena := ""
	nombre_archivo := "./grafo.dot"
	nombre_imagen := "grafo.jpg"
	if g.Principal != nil {
		cadena += "graph grafoDirigido{ \n rankdir=LR; \n node [shape=box]; layout=neato; \n nodo00[label=\"" + g.Principal.Valor + "\"]; \n"
		cadena += "node [shape = ellipse]; \n"
		cadena += g.retornarValoresMatriz()
		cadena += "\n}"
	}
	crearArchivo(nombre_archivo)
	escribirArchivo(cadena, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}

func (g *Grafo) retornarValoresMatriz() string {
	cadena := ""
	x := 0
	y := 1
	/*CREACION DE NODOS*/
	aux := g.Principal.Abajo //Filas
	aux1 := aux              //Columnas
	/*CREACION DE NODOS CON LABELS*/
	for aux != nil {
		for aux1 != nil {
			cadena += "nodo" + strconv.Itoa(x) + strconv.Itoa(y) + "[label=\"" + aux1.Valor + "\" ]; \n"
			aux1 = aux1.Siguiente
			x++
		}
		if aux != nil {
			aux = aux.Abajo
			aux1 = aux
		}
		x = 0
		y++
	}

	/*CONEXION DE NODOS*/
	x = 0
	y = 1
	aux = g.Principal.Abajo //Filas
	aux1 = aux              //Columnas
	/*CREACION DE NODOS CON LABELS*/
	for aux != nil {
		if aux1 != nil {
			cadena += "nodo00 -- "
			cadena += "nodo" + strconv.Itoa(x) + strconv.Itoa(y) + " -- "
			cadena += "nodo" + strconv.Itoa(x+1) + strconv.Itoa(y) + " -- "
			cadena += "nodo" + strconv.Itoa(x+2) + strconv.Itoa(y) + "[len=1.00]; \n"
		}
		if aux != nil {
			aux = aux.Abajo
			aux1 = aux
		}
		x = 0
		y++
	}

	return cadena
}
