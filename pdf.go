package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func GenerarPDF() {
	// Crear una instancia de PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Agregar una página
	pdf.AddPage()

	// Establecer la fuente
	pdf.SetFont("Arial", "B", 16)

	// Agregar un título
	pdf.Cell(40, 10, "Hola, este es un PDF creado en Go!")

	// Guardar el PDF en un archivo
	err := pdf.OutputFileAndClose("mi_pdf.pdf")
	if err != nil {
		fmt.Println("Error al generar el PDF:", err)
		return
	}

	fmt.Println("PDF generado exitosamente como 'mi_pdf.pdf'")
}
