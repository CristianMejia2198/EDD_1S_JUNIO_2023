package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Alumno struct {
	Carnet int    `json:"carnet"`
	Nombre string `json:"nombre"`
	Clase  *Clase `json:"clase"`
}

type Clase struct {
	Id           int    `json:"id"`
	Nombre_clase string `json:"nombre_clase"`
}

var Alumnado []Alumno
var Alumnos *Alumno = &Alumno{Carnet: 0, Nombre: "", Clase: &Clase{Id: 0, Nombre_clase: ""}}

func main() {
	r := mux.NewRouter()
	Alumnado = append(Alumnado, Alumno{Carnet: 201700918, Nombre: "Cristian", Clase: &Clase{Id: 772, Nombre_clase: "Estructuras"}})
	Alumnado = append(Alumnado, Alumno{Carnet: 201700919, Nombre: "Alberto", Clase: &Clase{Id: 777, Nombre_clase: "Compi 1"}})
	Alumnado = append(Alumnado, Alumno{Carnet: 201700920, Nombre: "Suy", Clase: &Clase{Id: 773, Nombre_clase: "Archivos"}})
	Alumnado = append(Alumnado, Alumno{Carnet: 201700921, Nombre: "Mejia", Clase: &Clase{Id: 774, Nombre_clase: "Bases 1"}})
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/memoria", HomeHandler1).Methods("GET")

	log.Fatal(http.ListenAndServe(":3001", r))
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(Alumnado)
}

func HomeHandler1(w http.ResponseWriter, req *http.Request) {
	agregar(2017000918, "Cristian Suy", 772, "Estructura de Datos")
	json.NewEncoder(w).Encode(&Alumnos)
}

func agregar(carnet int, nombre string, id int, clase string) {
	Alumnos.Carnet = carnet
	Alumnos.Nombre = nombre
	Alumnos.Clase.Id = id
	Alumnos.Clase.Nombre_clase = clase
}
