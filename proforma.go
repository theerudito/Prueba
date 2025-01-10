package main

type ProformaDetalle struct {
	ProformasID      int64   `json:"proformasid"`
	Emision          string  `json:"emision"`
	ProformasCodigo  string  `json:"proformas_codigo"`
	RazonSocial      string  `json:"razonsocial"`
	ProductoCodigo   string  `json:"productocodigo"`
	Detalle          string  `json:"detalle"`
	Cantidad         float64 `json:"cantidad"`
	Precio           float64 `json:"precio"`
	Descuento        float64 `json:"descuento"`
	Total            float64 `json:"total"`
	DescripcionCorta string  `json:"descripcioncorta"`
}
