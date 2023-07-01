package Facturas

type NodoBloque struct {
	Bloque    map[string]string
	Siguiente *NodoBloque
	Anterior  *NodoBloque
}
