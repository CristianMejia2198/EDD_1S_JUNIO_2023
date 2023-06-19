package estructuras

import "fmt"

type Arbol struct {
	Raiz *NodoArbol
}

func (a *Arbol) insertarNodo(raiz *NodoArbol, nuevoNodo *NodoArbol) *NodoArbol {
	if raiz == nil {
		raiz = nuevoNodo
	} else {
		if raiz.Valor > nuevoNodo.Valor {
			raiz.Izquierdo = a.insertarNodo(raiz.Izquierdo, nuevoNodo)
		} else {
			raiz.Derecho = a.insertarNodo(raiz.Derecho, nuevoNodo)
		}
	}
	return raiz
}

func (a *Arbol) InsertarElemento(valor int) {
	nuevoNodo := &NodoArbol{Valor: valor}
	a.Raiz = a.insertarNodo(a.Raiz, nuevoNodo)
}

func (a *Arbol) recorridoInorden(raiz *NodoArbol) {
	// Visita el lado izquierdo, luego pasa a la raiz, y por ultimo se va al lado derecho
	// izquierdo -> raiz -> derecho
	if raiz != nil {
		if raiz.Izquierdo != nil {
			a.recorridoInorden(raiz.Izquierdo)
			fmt.Print("->")
		}
		fmt.Print(" ", raiz.Valor, " ")
		if raiz.Derecho != nil {
			fmt.Print("->")
			a.recorridoInorden(raiz.Derecho)
		}
	}
}

func (a *Arbol) recorridoPreorden(raiz *NodoArbol) {
	//Raiz -> izquierdo -> derecho
	if raiz != nil {
		fmt.Print(raiz.Valor, " ")
		if raiz.Izquierdo != nil {
			fmt.Print("-> ")
			a.recorridoPreorden(raiz.Izquierdo)
		}
		if raiz.Derecho != nil {
			fmt.Print("-> ")
			a.recorridoPreorden(raiz.Derecho)
		}
	}
}

func (a *Arbol) recorridoPostorden(raiz *NodoArbol) {
	// izquierdo -> derecho -> raiz
	if raiz != nil {
		if raiz.Izquierdo != nil {
			a.recorridoPostorden(raiz.Izquierdo)
			fmt.Print("-> ")
		}
		if raiz.Derecho != nil {
			a.recorridoPostorden(raiz.Derecho)
			fmt.Print("-> ")
		}
		fmt.Print(raiz.Valor, " ")
	}
}

func (a *Arbol) GenerarRecorrido() {
	a.recorridoInorden(a.Raiz)
	fmt.Println("")
	a.recorridoPreorden(a.Raiz)
	fmt.Println("")
	a.recorridoPostorden(a.Raiz)
}
