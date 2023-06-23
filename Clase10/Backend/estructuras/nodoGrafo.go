package estructuras

type NodoMatrizDeAdyacencia struct {
	Siguiente *NodoMatrizDeAdyacencia
	Abajo     *NodoMatrizDeAdyacencia
	Valor     string
}

type EnvioMatriz struct {
	Padre   string
	Cliente string
	Imagen  string
	Filtros string
}
