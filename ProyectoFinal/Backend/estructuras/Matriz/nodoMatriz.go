package Matriz

type NodoMatriz struct {
	Siguiente *NodoMatriz
	Anterior  *NodoMatriz
	Abajo     *NodoMatriz
	Arriba    *NodoMatriz
	PosX      int
	PosY      int
	Color     string
}
