package main

import (
	"Clase10/estructuras"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var Matriz *estructuras.Grafo

func main() {
	Matriz = &estructuras.Grafo{Principal: nil}
	app := fiber.New()
	app.Use(cors.New())

	/*INICIO DE PETICIONES*/
	app.Post("/agregar-matriz", func(c *fiber.Ctx) error {
		var nuevoNodo estructuras.EnvioMatriz
		c.BodyParser(&nuevoNodo)
		Matriz.InsertarValores(nuevoNodo.Padre, nuevoNodo.Cliente, nuevoNodo.Imagen, nuevoNodo.Filtros)
		return c.JSON(&fiber.Map{
			"Data": nuevoNodo,
		})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"Data": Matriz,
		})
	})

	app.Get("/reporte-matriz", func(c *fiber.Ctx) error {
		Matriz.Reporte()
		return c.JSON(&fiber.Map{
			"Message": "OK",
		})
	})
	/*FIN DE PETICIONES*/
	app.Listen(":3001")
}
