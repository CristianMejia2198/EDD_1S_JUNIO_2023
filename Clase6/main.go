package main

import (
	"Clase6/estructuras"
)

func main() {
	matriz := &estructuras.Matriz{Raiz: &estructuras.NodoMatriz{PosX: -1, PosY: -1, Color: "RAIZ"}}
	matriz_csv := &estructuras.Matriz{Raiz: &estructuras.NodoMatriz{PosX: -1, PosY: -1, Color: "RAIZ"}}
	/*Primer Caso*/
	matriz.Insertar_Elemento(0, 0, "255-255-255")
	matriz.Insertar_Elemento(4, 4, "255-255-255")
	matriz.Insertar_Elemento(2, 2, "255-255-255")
	matriz.Insertar_Elemento(6, 6, "255-255-255")
	/*Segundo Caso*/
	matriz.Insertar_Elemento(0, 1, "255-255-255")
	matriz.Insertar_Elemento(6, 9, "255-255-255")
	/*Tercer Caso*/
	matriz.Insertar_Elemento(1, 6, "255-255-255")
	matriz.Insertar_Elemento(8, 9, "255-255-255")
	/*Cuarto Caso*/
	matriz.Insertar_Elemento(2, 6, "255-255-255")
	matriz.Insertar_Elemento(0, 4, "255-255-255")
	/*Especial*/
	matriz.Insertar_Elemento(0, 0, "0-0-0")
	//matriz.Reporte()
	imagen := "deadpool"
	//archivo := "body.csv"
	//matriz_csv.LeerArchivo("csv/" + imagen + "/" + archivo)
	//matriz_csv.Reporte()
	matriz_csv.LeerInicial("csv/"+imagen+"/inicial.csv", imagen)
	matriz_csv.GenerarImagen(imagen)
}
