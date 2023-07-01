package ColaPedidos

type NodoCola struct {
	Pedido    *PedidoCola
	Siguiente *NodoCola
}

type PedidoCola struct {
	Id_Cliente    int
	Nombre_Imagen string
}
