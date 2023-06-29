package main

import (
	"ProyectoFinal/estructuras/ArbolAVL"
	"ProyectoFinal/estructuras/Lista"
	"ProyectoFinal/estructuras/Matriz"
	"ProyectoFinal/estructuras/Peticiones"
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

func main() {
	/*INICIAR ESTRUCTURAS*/
	ListaEmpleado = &Lista.ListaDoble{Inicio: nil, Longitud: 0}
	ArbolPedidos = &ArbolAVL.Arbol{Raiz: nil}
	MatrizOriginal = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
	MatrizFiltro = &Matriz.Matriz{Raiz: &Matriz.NodoMatriz{PosX: -1, PosY: -1, Color: "Raiz"}}
	/**/
	imagen := "mario"
	MatrizOriginal.LeerInicial("csv/"+imagen+"/inicial.csv", imagen)
	MatrizOriginal.GenerarImagen(imagen)
	MatrizFiltro.LeerInicial("csv/"+imagen+"/inicial.csv", imagen)
	MatrizFiltro.Negativo("marioNegativo")
	/**/
	app := fiber.New()
	app.Use(cors.New())
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
		ArbolPedidos.Graficar()
		return c.JSON(&fiber.Map{
			"status": 200,
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

	app.Post("/aplicarfiltro", func(c *fiber.Ctx) error {
		var tipo Peticiones.PeticionFiltro
		c.BodyParser(&tipo)
		fmt.Println(tipo)
		if tipo.Tipo == 2 {
			MatrizFiltro.LeerInicial("csv/"+imagen+"/inicial.csv", imagen)
			MatrizFiltro.EscalaGrises("marioGrises")
		}
		return c.JSON(&fiber.Map{
			"status": 200,
		})
	})

	app.Listen(":3001")
}
