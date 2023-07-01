package Grafo

type NodoMatrizDeAdyacencia struct {
	Siguiente *NodoMatrizDeAdyacencia
	Abajo     *NodoMatrizDeAdyacencia
	Valor     string
}
