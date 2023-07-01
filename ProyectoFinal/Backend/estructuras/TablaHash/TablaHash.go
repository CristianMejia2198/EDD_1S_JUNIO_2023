package TablaHash

import (
	"fmt"
	"math"
	"strconv"
)

type TablaHash struct {
	Tabla       [30]NodoHash
	Capacidad   int // 5
	Utilizacion int // 0.0
}

func (t *TablaHash) calculoIndice(id_cliente int, multiplicador int) int {
	/*índice = (45*1526 + 202312345) % tamaño_tablahash*/
	indice := (18*id_cliente + 201700918*multiplicador) % t.Capacidad
	return indice
}

func (t *TablaHash) capacidad_Tabla() {
	auxCap := float64(t.Capacidad) * 0.6
	if t.Utilizacion > int(auxCap) {
		t.Capacidad = t.nuevaCapacidad()
		t.Utilizacion = 0
		t.reInsertar()
	}
}

func (t *TablaHash) reInsertar() {
	auxTabla := t.Tabla
	t.NewTablaHash()
	for i := 0; i < 30; i++ {
		if auxTabla[i].Llave != -1 {
			fmt.Println(auxTabla[i].Id_Cliente)
			t.Insertar(auxTabla[i].Id_Cliente, auxTabla[i].Id_Factura)
		}
	}
}

func (t *TablaHash) NewTablaHash() {
	for i := 0; i < 30; i++ {
		t.Tabla[i].Llave = -1
		t.Tabla[i].Id_Cliente = ""
		t.Tabla[i].Id_Factura = ""
	}
}

func (t *TablaHash) nuevaCapacidad() int {
	numero := t.Capacidad + 1
	for !t.isPrime(numero) {
		numero++
	}
	return numero
}

func (t *TablaHash) isPrime(numero int) bool {
	if numero <= 1 {
		return false
	}
	if numero == 2 {
		return true
	}
	if numero%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(numero))); i += 2 {
		if (numero % i) == 0 {
			return false
		}
	}
	return true
}

func (t *TablaHash) Insertar(id_cliente string, id_factura string) {
	numVar, _ := strconv.Atoi(id_cliente)
	indice := t.calculoIndice(numVar, 1)
	nuevoNodo := &NodoHash{Llave: indice, Id_Cliente: id_cliente, Id_Factura: id_factura}
	if indice < t.Capacidad {
		if t.Tabla[indice].Llave == -1 {
			t.Tabla[indice] = *nuevoNodo
			t.Utilizacion++
			t.capacidad_Tabla()
		} else {
			indice = t.calculoIndice(numVar, 2)
			if t.Tabla[indice].Llave == -1 {
				nuevoNodo.Llave = indice
				t.Tabla[indice] = *nuevoNodo
				t.Utilizacion++
				t.capacidad_Tabla()
				return
			}
			for i := indice; i < t.Capacidad; i++ {
				if t.Tabla[i].Llave == -1 {
					nuevoNodo.Llave = i
					t.Tabla[i] = *nuevoNodo
					t.Utilizacion++
					t.capacidad_Tabla()
					return
				}
			}
		}
	}
}
