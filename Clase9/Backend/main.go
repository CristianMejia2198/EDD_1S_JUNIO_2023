package main

import (
	"Clase9/estructuras"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RespuestaImagen struct {
	Imagenbase64 string
	Nombre       string
}

var arbol *estructuras.Arbol

func main() {
	arbol = &estructuras.Arbol{Raiz: nil}
	r := mux.NewRouter()
	/*Devolver algo*/
	r.HandleFunc("/", MostrarArbol).Methods("GET")
	/*Recibe un valor del frontend*/
	r.HandleFunc("/agregar-arbol", AgregarArbol).Methods("POST")
	r.HandleFunc("/reporte-arbol", MandarReporte).Methods("GET")
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

func MandarReporte(w http.ResponseWriter, req *http.Request) {
	arbol.Graficar()
	var imagen RespuestaImagen = RespuestaImagen{Nombre: "arbolAVL.jpg"}
	/*INICIO*/
	imageBytes, err := ioutil.ReadFile(imagen.Nombre)
	if err != nil {
		fmt.Fprintf(w, "Imagen No Valida")
		return
	}
	// Codifica los bytes de la imagen en base64
	imagen.Imagenbase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)

	/*data:image/jpg;base64,ABC*/
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(imagen)
}
