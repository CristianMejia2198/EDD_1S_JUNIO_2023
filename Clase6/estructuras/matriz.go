package estructuras

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Matriz struct {
	Raiz *NodoMatriz
}

func (m *Matriz) buscarC(x int) *NodoMatriz {
	aux := m.Raiz
	for aux != nil {
		if aux.PosX == x {
			return aux
		}
		aux = aux.Siguiente
	}
	return nil
}

func (m *Matriz) buscarF(y int) *NodoMatriz {
	aux := m.Raiz
	for aux != nil {
		if aux.PosY == y {
			return aux
		}
		aux = aux.Abajo
	}
	return nil
}

func (m *Matriz) insertarColumna(nuevoNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz {
	temp := nodoRaiz
	piv := false
	for { // while(true) [2][2][2][5][5] -> [N]
		if temp.PosX == nuevoNodo.PosX {
			temp.PosY = nuevoNodo.PosY
			temp.Color = nuevoNodo.Color
			return temp
		} else if temp.PosX > nuevoNodo.PosX {
			piv = true
			break
		}
		if temp.Siguiente != nil {
			temp = temp.Siguiente
		} else {
			break
		}
	}
	if piv {
		/*Asumir que nuevo = C1*/
		nuevoNodo.Siguiente = temp          // C2
		temp.Anterior.Siguiente = nuevoNodo // siguiente de raiz ahora es C1
		nuevoNodo.Anterior = temp.Anterior  // Anterior Raiz
		temp.Anterior = nuevoNodo           //
	} else {
		temp.Siguiente = nuevoNodo
		nuevoNodo.Anterior = temp
	}
	return nuevoNodo
}

func (m *Matriz) insertarFila(nuevoNodo *NodoMatriz, nodoRaiz *NodoMatriz) *NodoMatriz {
	temp := nodoRaiz
	piv := false
	for { //
		if temp.PosY == nuevoNodo.PosY {
			temp.PosX = nuevoNodo.PosX
			temp.Color = nuevoNodo.Color
			return temp
		} else if temp.PosY > nuevoNodo.PosY {
			piv = true
			break
		}
		if temp.Abajo != nil {
			temp = temp.Abajo
		} else {
			break
		}
	}
	if piv {
		/*Asumir que nuevo = C1*/
		nuevoNodo.Abajo = temp         // C2
		temp.Arriba.Abajo = nuevoNodo  // siguiente de raiz ahora es C1
		nuevoNodo.Arriba = temp.Arriba // Anterior Raiz
		temp.Arriba = nuevoNodo        //
	} else {
		temp.Abajo = nuevoNodo
		nuevoNodo.Arriba = temp
	}
	return nuevoNodo
}

func (m *Matriz) nuevaColumna(x int) *NodoMatriz {
	col := "C" + strconv.Itoa(x) // C1
	nuevoNodo := &NodoMatriz{PosX: x, PosY: -1, Color: col}
	columna := m.insertarColumna(nuevoNodo, m.Raiz)
	return columna
}

func (m *Matriz) nuevaFila(y int) *NodoMatriz {
	col := "F" + strconv.Itoa(y) // C1
	nuevoNodo := &NodoMatriz{PosX: -1, PosY: y, Color: col}
	fila := m.insertarFila(nuevoNodo, m.Raiz)
	return fila
}

func (m *Matriz) Insertar_Elemento(x int, y int, color string) {
	nuevoNodo := &NodoMatriz{PosX: x, PosY: y, Color: color}
	nodoColumna := m.buscarC(x)
	nodoFila := m.buscarF(y)
	/*
		1. Columna y Fila no Existe
		2. Columna si existe pero Fila no
		3. Fila si existe pero Columna no
		4. Ambos existen
	*/

	if nodoColumna == nil && nodoFila == nil {
		fmt.Println("Primer Caso")
		nodoColumna = m.nuevaColumna(x)
		nodoFila = m.nuevaFila(y)
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else if nodoColumna != nil && nodoFila == nil {
		fmt.Println("Segundo Caso")
		nodoFila = m.nuevaFila(y)
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else if nodoColumna == nil && nodoFila != nil {
		fmt.Println("Tercer Caso")
		nodoColumna = m.nuevaColumna(x)
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else if nodoColumna != nil && nodoFila != nil {
		fmt.Println("Cuarto Caso")
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else {
		fmt.Println("ERROR!!!!!!")
	}
}

func (m *Matriz) Reporte() {
	texto := ""
	nombre_archivo := "./matriz.dot"
	nombre_imagen := "matriz.jpg"
	aux1 := m.Raiz
	aux2 := m.Raiz
	aux3 := m.Raiz
	if aux1 != nil {
		texto = "digraph MatrizCapa{ \n node[shape=box] \n rankdir=UD; \n {rank=min; \n"
		/** Creacion de los nodos actuales */
		for aux1 != nil {
			texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + "[label=\"" + aux1.Color + "\" ,rankdir=LR,group=" + strconv.Itoa(aux1.PosX+1) + "]; \n"
			aux1 = aux1.Siguiente
		}
		texto += "}"
		for aux2 != nil {
			aux1 = aux2
			texto += "{rank=same; \n"
			for aux1 != nil {
				texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + "[label=\"" + aux1.Color + "\" ,group=" + strconv.Itoa(aux1.PosX+1) + "]; \n"
				aux1 = aux1.Siguiente
			}
			texto += "}"
			aux2 = aux2.Abajo
		}
		/** Conexiones entre los nodos de la matriz */
		aux2 = aux3
		for aux2 != nil {
			aux1 = aux2
			for aux1.Siguiente != nil {
				texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + " -> " + "nodo" + strconv.Itoa(aux1.Siguiente.PosX+1) + strconv.Itoa(aux1.Siguiente.PosY+1) + " [dir=both];\n"
				aux1 = aux1.Siguiente
			}
			aux2 = aux2.Abajo
		}
		aux2 = aux3
		for aux2 != nil {
			aux1 = aux2
			for aux1.Abajo != nil {
				texto += "nodo" + strconv.Itoa(aux1.PosX+1) + strconv.Itoa(aux1.PosY+1) + " -> " + "nodo" + strconv.Itoa(aux1.Abajo.PosX+1) + strconv.Itoa(aux1.Abajo.PosY+1) + " [dir=both];\n"
				aux1 = aux1.Abajo
			}
			aux2 = aux2.Siguiente
		}
		texto += "}"
	} else {
		texto = "No hay elementos en la matriz"
	}
	//fmt.Println(texto)
	crearArchivo(nombre_archivo)
	escribirArchivo(texto, nombre_archivo)
	ejecutar(nombre_imagen, nombre_archivo)
}

func (m *Matriz) LeerArchivo(ruta string) {
	//listaAux := &ListaCircular{Inicio: nil, Longitud: 0}
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	x := 0
	y := 0
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		for i := 0; i < len(linea); i++ {
			if linea[i] != "x" {
				m.Insertar_Elemento(x, y, linea[i])
			}
			x++
		}
		x = 0
		y++
	}
}
