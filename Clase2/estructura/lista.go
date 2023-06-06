package estructura

import "fmt"

type Lista struct {
	Inicio   *Nodo
	Longitud int
}

/*
this.inicio, this.longitud ---- JAVA
self.inicio, self.longitud ---- Python
this->inicio, this->longitud ---- C++
*/

func (l *Lista) estaVacia() bool {
	return l.Longitud == 0
}

/*
inicio = {1, 0x68984BA}
0x68984BA = {2, 0x7648420} -> aux
0x7648420 = {3, nil}
*/
func (l *Lista) Insertar(numero int) {
	if l.estaVacia() {
		l.Inicio = &Nodo{valor: numero, siguiente: nil}
		l.Longitud++
	} else {
		aux := l.Inicio
		for aux.siguiente != nil {
			aux = aux.siguiente
		}
		aux.siguiente = &Nodo{valor: numero, siguiente: nil}
		l.Longitud++
	}
}

func (l *Lista) Mostrar() {
	aux := l.Inicio

	for aux != nil {
		fmt.Println(aux.valor)
		aux = aux.siguiente
	}
}
