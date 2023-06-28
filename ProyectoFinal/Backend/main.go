package main

import (
	"ProyectoFinal/estructuras/Lista"
	"ProyectoFinal/estructuras/Peticiones"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var ListaEmpleado *Lista.ListaDoble

func main() {
	/*INICIAR ESTRUCTURAS*/
	ListaEmpleado = &Lista.ListaDoble{Inicio: nil, Longitud: 0}
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
		}
		return c.JSON(&fiber.Map{
			"status": "200",
		})
	})

	app.Post("/cargarempleados", func(c *fiber.Ctx) error {
		var nombreArchivo Peticiones.Archivo
		c.BodyParser(&nombreArchivo)
		fmt.Println(nombreArchivo)
		ListaEmpleado.LeerArchivo(nombreArchivo.Nombre)
		fmt.Println(ListaEmpleado.Inicio.Empleado.Id_Cliente)
		return c.JSON(&fiber.Map{
			"datos": ListaEmpleado,
		})
	})
	app.Listen(":3001")
}
