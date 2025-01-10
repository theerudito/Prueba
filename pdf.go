package main

import (
	"bytes"
	"encoding/base64"

	"github.com/jung-kurt/gofpdf"
)

func GenerarPDF_Bytes() []byte {
	// Crear una instancia de PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Agregar una página
	pdf.AddPage()

	// Establecer la fuente
	pdf.SetFont("Arial", "B", 16)

	// Agregar un título
	pdf.Cell(40, 10, "Hola, este es un PDF creado en Go!")

	// Crear un buffer para almacenar el PDF
	var buffer bytes.Buffer

	// Guardar el contenido del PDF en el buffer
	err := pdf.Output(&buffer)
	if err != nil {
		panic(err) // Maneja el error según sea necesario
	}

	// Retornar los bytes del PDF
	return buffer.Bytes()
}

func GenerarPDF_Base64() string {
	// Crear una instancia de PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Agregar una página
	pdf.AddPage()

	// Establecer la fuente
	pdf.SetFont("Arial", "B", 16)

	// Agregar un título
	pdf.Cell(40, 10, "Hola, este es un PDF creado en Go!")

	var buffer bytes.Buffer

	// Guardar el contenido del PDF en el buffer
	err := pdf.Output(&buffer)
	if err != nil {
		panic(err) // Maneja el error según sea necesario
	}

	// Convertir el buffer a Base64
	base64PDF := base64.StdEncoding.EncodeToString(buffer.Bytes())

	// Retornar el PDF en formato Base64
	return base64PDF

}
