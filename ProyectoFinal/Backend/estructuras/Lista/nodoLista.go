package Lista

type NodoLista struct {
	Empleado  *Empleado
	Siguiente *NodoLista
}

type Empleado struct {
	Id_Cliente string
	Password   string
}
