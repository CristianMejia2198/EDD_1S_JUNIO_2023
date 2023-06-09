package main

import (
	"Clase5/estructuras"
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
)

func main() {
	/* Representacion Rapida de la Lista Circular */
	var alumno []estructuras.Alumno = estructuras.LeerArchivo("Estudiante.csv")
	/* Representacion Rapida de la Lista Doblemente Enlazada*/
	imagen := [4]string{"mario", "estrellita", "cachorro", "letra"}

	colaActual := &estructuras.Cola{Primero: nil, Longitud: 0}
	pilaActual := &estructuras.Pila{Primero: nil, Longitud: 0}

	colaActual.LeerArchivo("EstudiantesCola.csv")

	opcion := 0
	salir := false

	for !salir {
		fmt.Println("1. Realizar Asignacion")
		fmt.Println("2. Reporte Cola")
		fmt.Println("3. Reporte Pila")
		fmt.Println("4. Salir")
		fmt.Print("Elige una opcion: ")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			existe := verificar(colaActual, alumno)
			if existe && colaActual.Primero != nil { /*Usuario ya esta registrado en lista circular y la Cola aun tiene elementos*/
				fmt.Println("El usuario actual es: ", colaActual.Primero.Alumno)
				imagenElegida := mostrarImagenes(imagen)
				pilaActual.Push(imagenElegida, colaActual.Primero.Alumno)
				colaActual.Descolar()
			} else if !existe && colaActual.Primero != nil {
				variable := asignarLista(colaActual, &alumno)
				fmt.Println(variable)
				imagenElegida := mostrarImagenes(imagen)
				pilaActual.Push(imagenElegida, colaActual.Primero.Alumno)
				colaActual.Descolar()
			} else if colaActual.Primero == nil {
				fmt.Println("Ya no hay alumnos por atender")
			}
		case 2:
			colaActual.Graficar()
			cmd := exec.Command("cmd", "/c", "start", "cola.jpg")
			err := cmd.Start()
			if err != nil {
				fmt.Println("No pude abrir imagen")
			}
		case 3:
			pilaActual.Graficar()
			cmd := exec.Command("cmd", "/c", "start", "pila.jpg")
			err := cmd.Start()
			if err != nil {
				fmt.Println("No pude abrir imagen")
			}
		case 4:
			salir = true
		}
		fmt.Print("\033[H\033[2J")
	}
	fmt.Println(alumno)
}

func verificar(colaActual *estructuras.Cola, alumnos []estructuras.Alumno) bool {
	aux := colaActual.Primero
	if aux != nil {
		if aux.Alumno.Carnet != "x" {
			for i := 0; i < len(alumnos); i++ {
				if aux.Alumno.Carnet == alumnos[i].Carnet {
					return true
				}
			}
		} else {
			return false
		}
	}
	return false
}

func mostrarImagenes(imagenes [4]string) string {
	fmt.Println("Imagenes Disponibles")
	opcion := 0
	for i := 0; i < 4; i++ {
		fmt.Println(strconv.Itoa(i+1), ". ", imagenes[i])
	}
	/*1,2,3,4*/
	fmt.Print("Elija una imagen: ")
	fmt.Scanln(&opcion)
	if (opcion - 1) < 4 {
		return imagenes[opcion-1]
	} else {
		mostrarImagenes(imagenes)
	}
	return ""
}

/*func asignarLista(colaActual *estructuras.Cola, listaActual *estructura.ListaCircular) string */
func asignarLista(colaActual *estructuras.Cola, alumnos *[]estructuras.Alumno) string {
	valor := (rand.Intn(1000)) + 202300000 // + 1000
	/*listaActual.Insertar(valor, colaActual.Primero.Alumno.Nombre)*/
	*alumnos = append(*alumnos, estructuras.Alumno{Carnet: strconv.Itoa(valor), Nombre: colaActual.Primero.Alumno.Nombre})
	cadena := "Se asigno el carnet: " + strconv.Itoa(valor) + " al estudiante " + colaActual.Primero.Alumno.Nombre
	colaActual.Primero.Alumno.Carnet = strconv.Itoa(valor)
	return cadena
}
