package estructuras

type NodoArbol struct {
	Izquierdo         *NodoArbol
	Derecho           *NodoArbol
	Valor             int
	Altura            int
	Factor_Equilibrio int
}
