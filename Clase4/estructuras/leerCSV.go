package estructuras

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func LeerArchivo(ruta string, listaAux *ListaCircular) {
	//listaAux := &ListaCircular{Inicio: nil, Longitud: 0}
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		listaAux.Insertar(linea[0], linea[1], linea[2])
	}
}
