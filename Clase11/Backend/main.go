package main

import (
	"Clase11/estructuras"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Persona struct {
	Identificador string
	Password      string
}

/*
{
	Identificador: "usuario1"
	Password: "1234"
}
*/

var Tablahash *estructuras.TablaHash

func main() {
	Tablahash = &estructuras.TablaHash{Capacidad: 5, Utilizacion: 0}
	Tablahash.NewTablaHash()
	app := fiber.New()
	app.Use(cors.New())
	// app.handlerFunc("/ruta",nombreFuncion).metodo("GET")
	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(&fiber.Map{
			"status": "ok",
		})
	})

	app.Post("/comprobar-usuario", func(c *fiber.Ctx) error {
		var usuario Persona
		c.BodyParser(&usuario)
		fmt.Println(usuario)
		if usuario.Identificador == "2017" && usuario.Password == "2017" {
			return c.JSON(&fiber.Map{
				"status": "OK",
			})
		}
		return c.JSON(&fiber.Map{
			"estado": "NO",
		})
	})

	app.Get("/obtener-tabla", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"Datos": Tablahash,
		})
	})

	/*8793, 4593, 7893, 2356*/
	app.Post("/agregar-tabla", func(c *fiber.Ctx) error {
		var nuevo estructuras.NodoHash
		c.BodyParser(&nuevo)
		Tablahash.Insertar(nuevo.Id_Cliente, nuevo.Id_Factura)
		return c.JSON(&fiber.Map{
			"status": "OK",
		})
	})

	app.Listen(":3001")
}
