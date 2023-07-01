package ColaPedidos

import (
	"fmt"
)

type Cola struct {
	Primero  *NodoCola
	Longitud int
}

func (c *Cola) Encolar(idCliente int, nombreimagen string) {
	nuevoPedido := &PedidoCola{Id_Cliente: idCliente, Nombre_Imagen: nombreimagen}
	if c.Longitud == 0 {
		nuevoNodo := &NodoCola{nuevoPedido, nil}
		c.Primero = nuevoNodo
		c.Longitud++
	} else {
		nuevoNodo := &NodoCola{nuevoPedido, nil}
		aux := c.Primero
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		aux.Siguiente = nuevoNodo
		c.Longitud++
	}
}

func (c *Cola) Descolar() {
	if c.Longitud == 0 {
		fmt.Println("No hay pedidos pendientes en la cola")
	} else {
		c.Primero = c.Primero.Siguiente
		c.Longitud--
	}
}
