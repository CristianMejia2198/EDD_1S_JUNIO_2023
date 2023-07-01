package main

import (
	"ProyectoFinal/estructuras/ArbolAVL"
	"ProyectoFinal/estructuras/ColaPedidos"
	"ProyectoFinal/estructuras/Facturas"
	"ProyectoFinal/estructuras/Grafo"
	"ProyectoFinal/estructuras/Lista"
	"ProyectoFinal/estructuras/Matriz"
	"ProyectoFinal/estructuras/Peticiones"
	"ProyectoFinal/estructuras/TablaHash"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var ListaEmpleado *Lista.ListaDoble
var ArbolPedidos *ArbolAVL.Arbol
var MatrizOriginal *Matriz.Matriz
var MatrizFiltro *Matriz.Matriz
var PedidosCola *ColaPedidos.Cola
var FacturasRealizadas *Facturas.BlockChain
var VerFacturasRealizadas *TablaHash.TablaHash
var FiltrosColocados string
var EmpleadoLogeado string
var GrafosEmpleados map[string]Grafo.Grafo

/*
				  project/controllers/account
                     ^                    \
                    /                      \
                   /                        \
                  /                         \/
         project/components/mux <--- project/controllers/base


{
	"8593": grafo
}

Falta
Ingresar al Grafo los datos de los filtros
Hacer el reporte del Blockchain
mostrar ambos reportes blockchain y grafo
*/

func main() {
	/*INICIAR ESTRUCTURAS*/
	ListaEmpleado = &Lista.ListaDoble{Inicio: nil, Longitud: 0}
	ArbolPedidos = &ArbolAVL.Arbol{Raiz: nil}
	MatrizOriginal = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
	MatrizFiltro = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
	PedidosCola = &ColaPedidos.Cola{Primero: nil, Longitud: 0}
	FacturasRealizadas = &Facturas.BlockChain{Bloques_Creados: 0}
	VerFacturasRealizadas = &TablaHash.TablaHash{Capacidad: 5, Utilizacion: 0}
	FiltrosColocados = ""
	EmpleadoLogeado = ""
	app := fiber.New()
	app.Use(cors.New())
	/*LOGIN*/
	app.Post("/login", func(c *fiber.Ctx) error {
		var usuario Peticiones.Login
		c.BodyParser(&usuario)
		if usuario.Username == "ADMIN_201700918" && usuario.Password == "admin" {
			return c.JSON(&fiber.Map{
				"status": "400",
			})
		} else {
			if ListaEmpleado.Inicio != nil {
				if ListaEmpleado.Buscar(usuario.Username, usuario.Password) {
					VerFacturasRealizadas = &TablaHash.TablaHash{Capacidad: 5, Utilizacion: 0}
					VerFacturasRealizadas.NewTablaHash()
					EmpleadoLogeado = usuario.Username
					return c.JSON(&fiber.Map{
						"status": "200",
					})
				}
			}
		}
		return c.JSON(&fiber.Map{
			"status": "404",
		})
	})

	app.Post("/cargarempleados", func(c *fiber.Ctx) error {
		var nombreArchivo Peticiones.Archivo
		c.BodyParser(&nombreArchivo)
		fmt.Println(nombreArchivo)
		ListaEmpleado.LeerArchivo(nombreArchivo.Nombre)
		return c.JSON(&fiber.Map{
			"status": 200,
		})
	})

	app.Post("/cargarpedidos", func(c *fiber.Ctx) error {
		var pedidosNuevos Peticiones.ArbolPeticion
		c.BodyParser(&pedidosNuevos)
		for i := 0; i < len(pedidosNuevos.Pedidos); i++ {
			ArbolPedidos.InsertarElemento(pedidosNuevos.Pedidos[i].Id_Cliente, pedidosNuevos.Pedidos[i].Nombre_Imagen)
		}
		ArbolPedidos.RecorridoInorden(ArbolPedidos.Raiz, PedidosCola)
		return c.JSON(&fiber.Map{
			"status": 200,
			"cola":   PedidosCola,
		})
	})

	app.Get("/reporte-arbol", func(c *fiber.Ctx) error {
		var imagen Peticiones.RespuestaImagen = Peticiones.RespuestaImagen{Nombre: "Reporte/arbolAVL.jpg"}
		/*INICIO*/
		imageBytes, err := ioutil.ReadFile(imagen.Nombre)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status": 404,
			})
		}
		// Codifica los bytes de la imagen en base64
		imagen.Imagenbase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)
		return c.JSON(&fiber.Map{
			"status": 200,
			"imagen": imagen,
		})
	})

	app.Get("/reporte-grafo", func(c *fiber.Ctx) error {
		var imagen Peticiones.RespuestaImagen = Peticiones.RespuestaImagen{Nombre: "Reporte/grafo.jpg"}
		/*INICIO*/
		imageBytes, err := ioutil.ReadFile(imagen.Nombre)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status": 404,
			})
		}
		// Codifica los bytes de la imagen en base64
		imagen.Imagenbase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)
		return c.JSON(&fiber.Map{
			"status": 200,
			"imagen": imagen,
		})
	})

	app.Get("/reporte-bloque", func(c *fiber.Ctx) error {
		var imagen Peticiones.RespuestaImagen = Peticiones.RespuestaImagen{Nombre: "Reporte/bloque.jpg"}
		/*INICIO*/
		imageBytes, err := ioutil.ReadFile(imagen.Nombre)
		if err != nil {
			return c.JSON(&fiber.Map{
				"status": 404,
			})
		}
		// Codifica los bytes de la imagen en base64
		imagen.Imagenbase64 = "data:image/jpg;base64," + base64.StdEncoding.EncodeToString(imageBytes)
		return c.JSON(&fiber.Map{
			"status": 200,
			"imagen": imagen,
		})
	})

	app.Post("/aplicarfiltro", func(c *fiber.Ctx) error {
		var tipo Peticiones.PeticionFiltro
		c.BodyParser(&tipo)
		fmt.Println(tipo)
		tipo.NombreImagen = PedidosCola.Primero.Pedido.Nombre_Imagen
		switch tipo.Tipo {
		case 1:
			MatrizFiltro.LeerInicial("csv/"+tipo.NombreImagen+"/inicial.csv", tipo.NombreImagen)
			MatrizFiltro.Negativo(tipo.NombreImagen + "Negativo")
			FiltrosColocados += "Negativo, "
		case 2:
			MatrizFiltro.LeerInicial("csv/"+tipo.NombreImagen+"/inicial.csv", tipo.NombreImagen)
			MatrizFiltro.EscalaGrises(tipo.NombreImagen + "Grises")
			FiltrosColocados += "Grises, "
		case 3:
			MatrizFiltro.LeerInicial("csv/"+tipo.NombreImagen+"/inicial.csv", tipo.NombreImagen)
			MatrizFiltro.RotacionX()
			MatrizFiltro.GenerarImagen(tipo.NombreImagen + "RX")
			FiltrosColocados += "Eje X, "
		case 4:
			MatrizFiltro.LeerInicial("csv/"+tipo.NombreImagen+"/inicial.csv", tipo.NombreImagen)
			MatrizFiltro.RotacionY()
			MatrizFiltro.GenerarImagen(tipo.NombreImagen + "RY")
			FiltrosColocados += "Eje Y, "
		case 5:
			MatrizFiltro.LeerInicial("csv/"+tipo.NombreImagen+"/inicial.csv", tipo.NombreImagen)
			MatrizFiltro.RotacionDoble()
			MatrizFiltro.GenerarImagen(tipo.NombreImagen + "RDoble")
			FiltrosColocados += "Doble,  "
		}
		return c.JSON(&fiber.Map{
			"status": 200,
		})
	})

	app.Get("/obtenerPedido", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"datos": PedidosCola.Primero.Pedido,
		})
	})

	app.Post("/generarfactura", func(c *fiber.Ctx) error {
		var nuevoBloque Peticiones.BloquePeticion
		c.BodyParser(&nuevoBloque)
		FacturasRealizadas.InsertarBloque(nuevoBloque.Timestamp, nuevoBloque.Biller, nuevoBloque.Customer, nuevoBloque.Payment)
		/*Ingresar al grafo, tomar los valores de nuevoBloque.Biller, nuevoBloque.Customer, PedidosCola.Primero.Pedido.Nombre_Imagen,Filtros_colocados */
		PedidosCola.Descolar()
		VerFacturasRealizadas.NewTablaHash()
		FacturasRealizadas.InsertarTabla(VerFacturasRealizadas, EmpleadoLogeado)
		MatrizOriginal = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
		MatrizFiltro = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
		return c.JSON(&fiber.Map{
			"datos": FacturasRealizadas.Bloques_Creados,
		})
	})

	app.Get("/facturaempleado", func(c *fiber.Ctx) error {

		return c.JSON(&fiber.Map{
			"status":  200,
			"factura": VerFacturasRealizadas.Tabla,
		})
	})

	app.Listen(":3001")
}
