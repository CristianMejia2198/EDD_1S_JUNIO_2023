package estructuras

import "fmt"

type ListaCircular struct {
	Inicio   *nodoLista
	Longitud int
}

func (l *ListaCircular) Insertar(carnet string, nombre string, password string) {
	nuevoAlumno := &Alumno{Carnet: carnet, Nombre: nombre, Password: password}
	if l.Longitud == 0 {
		l.Inicio = &nodoLista{alumno: nuevoAlumno, siguiente: nil}
		l.Inicio.siguiente = l.Inicio
		l.Longitud++
	} else {
		if l.Longitud == 1 {
			l.Inicio.siguiente = &nodoLista{alumno: nuevoAlumno, siguiente: l.Inicio}
			l.Longitud++
		} else {
			aux := l.Inicio
			for i := 0; i < l.Longitud-1; i++ {
				aux = aux.siguiente
			}
			aux.siguiente = &nodoLista{alumno: nuevoAlumno, siguiente: l.Inicio}
			l.Longitud++
		}
	}
}

func (l *ListaCircular) Mostrar() {
	aux := l.Inicio
	for i := 0; i < l.Longitud; i++ {
		fmt.Println("Nombre: ", aux.alumno.Nombre, " Carnet: ", aux.alumno.Carnet, " Password: ", aux.alumno.Password)
		aux = aux.siguiente
	}
}
