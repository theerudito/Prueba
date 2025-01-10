package main

import (
	"database/sql"
	"log"
)

func ConsultaProformas(db *sql.DB) []ProformaDetalle {

	query := `
		SELECT 
			proforma.proformasid AS proformasid,	
			proforma.emision AS emision,	
			proforma.proformas_codigo AS proformas_codigo,	
			cliente.razonsocial AS razonsocial,	
			producto.productocodigo AS productocodigo,	
			producto.descripcion AS detalle,	
			detalles.cantidaddigitada AS cantidad,	
			detalles.preciovisible AS precio,	
			detalles.descuento AS descuento,	
			(detalles.cantidaddigitada * detalles.preciovisible) AS total,	
			medida.descripcioncorta AS descripcioncorta
		FROM 
			proformas AS proforma
		JOIN 
			clientes AS cliente ON proforma.clientesid = cliente.clientesid
		JOIN 
			proformas_detalles AS detalles ON detalles.proformasid = proforma.proformasid
		JOIN 
			productos AS producto ON producto.productosid = detalles.productosid
		JOIN 
			medidas medida ON medida.medidasid = detalles.medidasid
	`

	// Resultado de la consulta
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error al ejecutar la consulta:", err)
	}
	defer rows.Close()

	// Iterar sobre los resultados
	var resultados []ProformaDetalle
	for rows.Next() {
		var detalle ProformaDetalle
		err := rows.Scan(
			&detalle.ProformasID,
			&detalle.Emision,
			&detalle.ProformasCodigo,
			&detalle.RazonSocial,
			&detalle.ProductoCodigo,
			&detalle.Detalle,
			&detalle.Cantidad,
			&detalle.Precio,
			&detalle.Descuento,
			&detalle.Total,
			&detalle.DescripcionCorta,
		)
		if err != nil {
			log.Fatal("Error al escanear los resultados:", err)
		}
		resultados = append(resultados, detalle)
	}

	return resultados
}
