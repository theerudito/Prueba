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
	pdf := gofpdf.New("L", "mm", "A4", "")

	// Agregar una página
	pdf.AddPage()

	// Establecer la fuente para el encabezado
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(200, 10, "Reporte de Proformas")
	pdf.Ln(8) // Salto de línea

	// Encabezado de la tabla
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(20, 5, "Emision", "", 0, "C", false, 0, "")
	pdf.CellFormat(25, 5, "Codigo", "", 0, "C", false, 0, "")
	pdf.CellFormat(80, 5, "Cliente", "", 0, "C", false, 0, "")
	pdf.CellFormat(50, 5, "Producto", "", 0, "C", false, 0, "")
	pdf.CellFormat(20, 5, "Medida", "", 0, "C", false, 0, "")
	pdf.CellFormat(15, 5, "Cant.", "", 0, "C", false, 0, "")
	pdf.CellFormat(15, 5, "Precio", "", 0, "C", false, 0, "")
	pdf.CellFormat(15, 5, "Desc", "", 0, "C", false, 0, "")
	pdf.CellFormat(15, 5, "Subtotal", "", 0, "C", false, 0, "")
	pdf.CellFormat(15, 5, "Total", "", 0, "C", false, 0, "")
	pdf.Ln(-1)

	// Contenido de la tabla
	pdf.SetFont("Arial", "", 8)
	var totalGeneral float64
	for _, detalle := range detalle {
		pdf.CellFormat(20, 5, detalle.Emision, "", 0, "C", false, 0, "")
		pdf.CellFormat(25, 5, detalle.ProformasCodigo, "", 0, "L", false, 0, "")
		pdf.CellFormat(80, 5, detalle.RazonSocial, "", 0, "L", false, 0, "")
		pdf.CellFormat(50, 5, detalle.Detalle, "", 0, "L", false, 0, "")
		pdf.CellFormat(20, 5, detalle.DescripcionCorta, "", 0, "C", false, 0, "")
		pdf.CellFormat(15, 5, fmt.Sprintf("%.2f", detalle.Cantidad), "", 0, "C", false, 0, "")
		pdf.CellFormat(15, 5, fmt.Sprintf("%.2f", detalle.Precio), "", 0, "R", false, 0, "")
		pdf.CellFormat(15, 5, fmt.Sprintf("%.2f", detalle.Descuento), "", 0, "R", false, 0, "")
		pdf.CellFormat(15, 5, fmt.Sprintf("%.2f", detalle.Total), "", 0, "R", false, 0, "")
		pdf.CellFormat(15, 5, fmt.Sprintf("%.2f", detalle.Total), "", 0, "R", false, 0, "")
		pdf.Ln(-1)

		totalGeneral += detalle.Total
	}

	// Total General
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(250, 10, "Total General", "", 0, "R", false, 0, "")
	pdf.CellFormat(20, 10, fmt.Sprintf("%.2f", totalGeneral), "", 0, "R", false, 0, "")
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
