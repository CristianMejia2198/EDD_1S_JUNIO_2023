package Peticiones

type Login struct {
	Username string
	Password string
}

type Archivo struct {
	Nombre string
}

type Pedido struct {
	Id_Cliente    int    `json:"id_cliente"`
	Nombre_Imagen string `json:"imagen"`
}

type ArbolPeticion struct {
	Pedidos []Pedido
}

type RespuestaImagen struct {
	Imagenbase64 string
	Nombre       string
}

type PeticionFiltro struct {
	Tipo         int
	NombreImagen string
}

type BloquePeticion struct {
	Timestamp string
	Biller    string
	Customer  string
	Payment   string
}

type Empleado struct {
	IdEmpleado string
}
