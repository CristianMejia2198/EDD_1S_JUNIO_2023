package main

import (
	"Clase12/estructuras"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var ListaBloque *estructuras.BlockChain

func main() {
	ListaBloque = &estructuras.BlockChain{Bloques_Creados: 0}
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/agregar-bloque", func(c *fiber.Ctx) error {
		var nuevoNodo estructuras.NodoBloquePeticion
		c.BodyParser(&nuevoNodo)
		ListaBloque.InsertarBloque(nuevoNodo.Timestamp, nuevoNodo.Biller, nuevoNodo.Customer, nuevoNodo.Payment)
		return c.JSON(&fiber.Map{
			"Data": nuevoNodo,
		})
	})

	app.Get("/tablafacturas", func(c *fiber.Ctx) error {
		factura := ListaBloque.ArregloFacturas()
		return c.JSON(&fiber.Map{
			"message": "OK",
			"data":    factura,
		})
	})

	app.Listen(":3001")
}
