package estructuras

type NodoCola struct {
	Alumno    *Alumno /* Representacion de Cliente */
	Siguiente *NodoCola
}

type NodoPila struct {
	Alumno    *Alumno /*Representacion de Cliente*/
	Color     string  /*Representacion de Imagen*/
	Siguiente *NodoPila
}
