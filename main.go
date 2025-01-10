package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db := ConexionDB()
	defer db.Close()

	// Realizar la consulta
	consulta := ConsultaProformas(db)

	bytesPDF := GenerarPDF(consulta)

	base64PDF := GenerarPDF(consulta)

	app := fiber.New()
	// Ruta para retornar bytes
	app.Get("/bytes", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/pdf")
		c.Set("Content-Disposition", "inline; filename=\"document.pdf\"")
		return c.Send(bytesPDF.Bytes)
	})

	//Ruta para retornar base64
	app.Get("/base64", func(c *fiber.Ctx) error {

		return c.SendString(base64PDF.Base64)
	})

	log.Fatal(app.Listen(":3000"))
}
