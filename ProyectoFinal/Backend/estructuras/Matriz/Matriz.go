package Matriz

import (
	"ProyectoFinal/estructuras/GenerarArchivos"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Matriz struct {
	Raiz        *NodoMatriz
	ImageWidth  int
	ImageHeight int
	PixelWidth  int
	PixelHeight int
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
		nodoColumna = m.nuevaColumna(x)
		nodoFila = m.nuevaFila(y)
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else if nodoColumna != nil && nodoFila == nil {
		nodoFila = m.nuevaFila(y)
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else if nodoColumna == nil && nodoFila != nil {
		nodoColumna = m.nuevaColumna(x)
		nuevoNodo = m.insertarColumna(nuevoNodo, nodoFila)
		nuevoNodo = m.insertarFila(nuevoNodo, nodoColumna)
	} else if nodoColumna != nil && nodoFila != nil {
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
	GenerarArchivos.CrearArchivo(nombre_archivo)
	GenerarArchivos.EscribirArchivo(texto, nombre_archivo)
	GenerarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
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

func (m *Matriz) LeerInicial(ruta string, imagen string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		if linea[0] == "0" {
			m.leerConfig("csv/" + imagen + "/" + linea[1]) /*csv/mario/config.csv*/
		} else {
			m.LeerArchivo("csv/" + imagen + "/" + linea[1])
		}
	}
}

func (m *Matriz) leerConfig(ruta string) {
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		numero, _ := strconv.Atoi(linea[1])
		if linea[0] == "image_width" {
			m.ImageWidth = numero
		} else if linea[0] == "image_height" {
			m.ImageHeight = numero
		} else if linea[0] == "pixel_width" {
			m.PixelWidth = numero
		} else if linea[0] == "pixel_height" {
			m.PixelHeight = numero
		}
	}
}

func (m *Matriz) GenerarImagen(nombre_imagen string) {
	archivoCSS := "ImagenFinal/" + nombre_imagen + ".css" // csv/mario/mario.css
	contenidoCSS := "body{\n background: #333333; \n height: 100vh; \n display: flex; \n justify-content: center; \n align-items: center; \n } \n"
	contenidoCSS += ".canvas{ \n width: " + strconv.Itoa(m.ImageWidth*m.PixelWidth) + "px; \n"
	contenidoCSS += "height: " + strconv.Itoa(m.ImageHeight*m.PixelHeight) + "px; \n }"
	contenidoCSS += ".pixel{ \n width: " + strconv.Itoa(m.PixelWidth) + "px; \n"
	contenidoCSS += "height: " + strconv.Itoa(m.PixelHeight) + "px; \n float: left; \n } \n"
	x_pixel := 0
	x := 1
	auxFila := m.Raiz.Abajo
	auxColumna := auxFila.Siguiente

	//* Nueva Version*//
	for i := 0; i < m.ImageHeight; i++ {
		for j := 0; j < m.ImageWidth; j++ {
			if auxColumna != nil {
				if auxColumna.PosX == x_pixel {
					contenidoCSS += ".pixel:nth-child(" + strconv.Itoa(x) + ") { background: rgb(" + strings.ReplaceAll(auxColumna.Color, "-", ",") + "); }\n"
					auxColumna = auxColumna.Siguiente
				}
				x_pixel++
			}
			x++
		}
		x_pixel = 0
		auxFila = auxFila.Abajo
		if auxFila != nil {
			auxColumna = auxFila.Siguiente
		}
	}

	/*FIN*/
	m.generarHTML(nombre_imagen)
	GenerarArchivos.CrearArchivo(archivoCSS)
	GenerarArchivos.EscribirArchivo(contenidoCSS, archivoCSS)
}

func (m *Matriz) generarHTML(nombre_imagen string) {
	archivoHTML := "ImagenFinal/" + nombre_imagen + ".html"
	contenidoHTML := "<!DOCTYPE html> \n <html> \n <head> \n <link rel=\"stylesheet\"  href=\""
	contenidoHTML += nombre_imagen + ".css"
	contenidoHTML += "\" > \n </head> \n <body> \n <div class=\"canvas\"> \n"
	for i := 0; i < m.ImageHeight; i++ {
		for j := 0; j < m.ImageWidth; j++ {
			contenidoHTML += "    <div class=\"pixel\"></div> \n"
		}
	}
	contenidoHTML += "</div> \n </body> \n </html> \n"
	GenerarArchivos.CrearArchivo(archivoHTML)
	GenerarArchivos.EscribirArchivo(contenidoHTML, archivoHTML)
}

func (m *Matriz) EscalaGrises(nombre_imagen string) {
	archivoCSS := "ImagenFinal/" + nombre_imagen + ".css" // csv/mario/mario.css
	contenidoCSS := "body{\n background: #333333; \n height: 100vh; \n display: flex; \n justify-content: center; \n align-items: center; \n } \n"
	contenidoCSS += ".canvas{ \n width: " + strconv.Itoa(m.ImageWidth*m.PixelWidth) + "px; \n"
	contenidoCSS += "height: " + strconv.Itoa(m.ImageHeight*m.PixelHeight) + "px; \n }"
	contenidoCSS += ".pixel{ \n width: " + strconv.Itoa(m.PixelWidth) + "px; \n"
	contenidoCSS += "height: " + strconv.Itoa(m.PixelHeight) + "px; \n float: left; \n } \n"
	x_pixel := 0
	x := 1
	auxFila := m.Raiz.Abajo
	auxColumna := auxFila.Siguiente
	// 0.299 ∙ Rojo + 0.587 ∙ Verde + 0.114 ∙ Azul
	//* Nueva Version*//
	for i := 0; i < m.ImageHeight; i++ {
		for j := 0; j < m.ImageWidth; j++ {
			if auxColumna != nil {
				if auxColumna.PosX == x_pixel {
					/*PROCEDIMIENTO FILTROS*/
					colores := strings.Split(auxColumna.Color, "-")
					r, _ := strconv.Atoi(colores[0])
					b, _ := strconv.Atoi(colores[1])
					g, _ := strconv.Atoi(colores[2])
					colorFinal := 0.299*float64(r) + 0.587*float64(b) + 0.114*float64(g) // 125-125-125
					auxColumna.Color = strconv.Itoa(int(colorFinal)) + "," + strconv.Itoa(int(colorFinal)) + "," + strconv.Itoa(int(colorFinal))
					/**/
					contenidoCSS += ".pixel:nth-child(" + strconv.Itoa(x) + ") { background: rgb(" + auxColumna.Color + "); }\n"
					auxColumna = auxColumna.Siguiente
				}
				x_pixel++
			}
			x++
		}
		x_pixel = 0
		auxFila = auxFila.Abajo
		if auxFila != nil {
			auxColumna = auxFila.Siguiente
		}
	}

	/*FIN*/
	m.generarHTML(nombre_imagen)
	GenerarArchivos.CrearArchivo(archivoCSS)
	GenerarArchivos.EscribirArchivo(contenidoCSS, archivoCSS)
}

func (m *Matriz) Negativo(nombre_imagen string) {
	archivoCSS := "ImagenFinal/" + nombre_imagen + ".css" // csv/mario/mario.css
	contenidoCSS := "body{\n background: #333333; \n height: 100vh; \n display: flex; \n justify-content: center; \n align-items: center; \n } \n"
	contenidoCSS += ".canvas{ \n width: " + strconv.Itoa(m.ImageWidth*m.PixelWidth) + "px; \n"
	contenidoCSS += "height: " + strconv.Itoa(m.ImageHeight*m.PixelHeight) + "px; \n }"
	contenidoCSS += ".pixel{ \n width: " + strconv.Itoa(m.PixelWidth) + "px; \n"
	contenidoCSS += "height: " + strconv.Itoa(m.PixelHeight) + "px; \n float: left; \n } \n"
	x_pixel := 0
	x := 1
	auxFila := m.Raiz.Abajo
	auxColumna := auxFila.Siguiente
	// 0.299 ∙ Rojo + 0.587 ∙ Verde + 0.114 ∙ Azul
	//* Nueva Version*//
	for i := 0; i < m.ImageHeight; i++ {
		for j := 0; j < m.ImageWidth; j++ {
			if auxColumna != nil {
				if auxColumna.PosX == x_pixel {
					/*PROCEDIMIENTO FILTROS*/
					colores := strings.Split(auxColumna.Color, "-")
					r, _ := strconv.Atoi(colores[0])
					b, _ := strconv.Atoi(colores[1])
					g, _ := strconv.Atoi(colores[2])
					if r >= 255 {
						r = 0
					} else {
						r = 255 - r
					}
					if g >= 255 {
						g = 0
					} else {
						g = 255 - g
					}
					if b >= 255 {
						b = 0
					} else {
						b = 255 - b
					}
					auxColumna.Color = strconv.Itoa(r) + "," + strconv.Itoa(g) + "," + strconv.Itoa(b)
					/**/
					contenidoCSS += ".pixel:nth-child(" + strconv.Itoa(x) + ") { background: rgb(" + auxColumna.Color + "); }\n"
					auxColumna = auxColumna.Siguiente
				}
				x_pixel++
			}
			x++
		}
		x_pixel = 0
		auxFila = auxFila.Abajo
		if auxFila != nil {
			auxColumna = auxFila.Siguiente
		}
	}

	/*FIN*/
	m.generarHTML(nombre_imagen)
	GenerarArchivos.CrearArchivo(archivoCSS)
	GenerarArchivos.EscribirArchivo(contenidoCSS, archivoCSS)
}
