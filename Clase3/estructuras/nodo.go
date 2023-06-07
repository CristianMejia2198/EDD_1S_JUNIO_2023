package estructuras

type NodoLista struct {
	alumno    *Alumno
	siguiente *NodoLista
	anterior  *NodoLista
}
