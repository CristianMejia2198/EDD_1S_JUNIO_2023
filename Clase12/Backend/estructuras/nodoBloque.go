package estructuras

type NodoBloque struct {
	Bloque    map[string]string
	Siguiente *NodoBloque
	Anterior  *NodoBloque
}

type NodoBloquePeticion struct {
	Timestamp string
	Biller    string
	Customer  string
	Payment   string
}

type RespuestaBloque struct {
	Id      string
	Factura string
}

/*
map[string]string
{
	"nombre": "nombre"
}

map[int]string
{
	2: "nombre"
}

map[int]int
{
	2: 15
}

*/
