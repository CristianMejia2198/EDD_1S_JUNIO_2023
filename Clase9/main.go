package main

import (
	"Clase9/estructuras"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var arbol *estructuras.Arbol

func main() {
	arbol = &estructuras.Arbol{Raiz: nil}
	r := mux.NewRouter()
	/*Devolver algo*/
	r.HandleFunc("/", MostrarArbol).Methods("GET")
	/*Recibe un valor del frontend*/
	r.HandleFunc("/agregar-arbol", AgregarArbol).Methods("POST")
	log.Fatal(http.ListenAndServe(":3001", r))
}

func MostrarArbol(w http.ResponseWriter, req *http.Request) {
	/*Esto nos verifica que le estamos enviando al servidor una respuesta de tipo JSON*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&arbol)
}

func AgregarArbol(w http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	var nuevoNodo estructuras.NodoArbol
	if err != nil {
		fmt.Fprintf(w, "No valido")
	}
	json.Unmarshal(reqBody, &nuevoNodo)
	arbol.InsertarElemento(nuevoNodo.Valor)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoNodo)
}
