package main

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

type PDFResult struct {
	Bytes  []byte
	Base64 string
}

func GenerarPDF(detalle []ProformaDetalle) PDFResult {
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Agregar una página
	pdf.AddPage()

	// Establecer la fuente para el encabezado
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "Reporte de Proformas")
	pdf.Ln(12) // Salto de línea

	// Encabezado de la tabla
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(20, 10, "Emision", "1", 0, "C", false, 0, "")
	pdf.CellFormat(25, 10, "Codigo", "1", 0, "C", false, 0, "")
	pdf.CellFormat(50, 10, "Cliente", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 10, "Producto", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 10, "Medida", "1", 0, "C", false, 0, "")
	pdf.CellFormat(15, 10, "Cant.", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 10, "Precio", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 10, "Descuento", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 10, "Subtotal", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

	// Contenido de la tabla
	pdf.SetFont("Arial", "", 10)
	var totalGeneral float64
	for _, detalle := range detalle {
		pdf.CellFormat(20, 10, detalle.Emision, "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 10, detalle.ProformasCodigo, "1", 0, "L", false, 0, "")
		pdf.CellFormat(50, 10, detalle.RazonSocial, "1", 0, "L", false, 0, "")
		pdf.CellFormat(30, 10, detalle.ProductoCodigo, "1", 0, "L", false, 0, "")
		pdf.CellFormat(20, 10, detalle.DescripcionCorta, "1", 0, "C", false, 0, "")
		pdf.CellFormat(15, 10, fmt.Sprintf("%.2f", detalle.Cantidad), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 10, fmt.Sprintf("%.2f", detalle.Precio), "1", 0, "R", false, 0, "")
		pdf.CellFormat(20, 10, fmt.Sprintf("%.2f", detalle.Descuento), "1", 0, "R", false, 0, "")
		pdf.CellFormat(20, 10, fmt.Sprintf("%.2f", detalle.Total), "1", 0, "R", false, 0, "")
		pdf.Ln(-1)

		totalGeneral += detalle.Total
	}

	// Total General
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(200, 10, "Total General", "1", 0, "R", false, 0, "")
	pdf.CellFormat(20, 10, fmt.Sprintf("%.2f", totalGeneral), "1", 0, "R", false, 0, "")
	pdf.Ln(-1)

	var buffer bytes.Buffer

	// Crear un buffer para almacenar el PDF
	err := pdf.Output(&buffer)
	if err != nil {
		panic(err)
	}

	// Obtener bytes y codificar en base64
	bytesPDF := buffer.Bytes()
	base64PDF := base64.StdEncoding.EncodeToString(bytesPDF)

	return PDFResult{
		Bytes:  bytesPDF,
		Base64: base64PDF,
	}
}
