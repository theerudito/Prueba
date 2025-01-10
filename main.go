package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	valor1 := GenerarPDF_Bytes()
	valor2 := GenerarPDF_Base64()

	//println("PDF generado con tamaño: %d bytes\n", len(valor1))
	//ConexionDB()

	app := fiber.New()
	// Ruta para retornar bytes
	app.Get("/bytes", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/pdf")
		c.Set("Content-Disposition", "inline; filename=\"document.pdf\"")
		return c.Send(valor1)
	})

	// Ruta para retornar base64
	app.Get("/base64", func(c *fiber.Ctx) error {

		return c.SendString(valor2)
	})

	log.Fatal(app.Listen(":3000"))
}
