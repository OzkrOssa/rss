package models

//API RESPONSE
/*
{
	"success": true,
	"data": {
		"resultado": "ok",
		"info": [
			{
				"id_contrato": "",
				"nro_contrato": "",
				"nombre": "",
				"apellido": "",
				"cedula": "",
				"inicial_doc": "",
				"direccion": "",
				"telefono": "",
				"telf_casa": "",
				"telf_adic": "",
				"status_contrato": "",
				"suscripcion": "",
				"servicio": [
					{
						"tipo_servicio": "",
						"nombre_servicio": "",
						"status": "",
						"monto": ""
					},
					{
						"tipo_servicio": "",
						"nombre_servicio": "",
						"status": "",
						"monto": ""
					}
				]
			}
		]
	}
}
*/

type Response struct {
	Success bool     `json:"success"`
	Data    DataInfo `json:"data"`
}

type DataInfo struct {
	Resultado string    `json:"resultado"`
	Info      []*Client `json:"info"`
}
type Client struct {
	IDContrato     string `json:"id_contrato"`
	NroContrato    string `json:"nro_contrato"`
	Nombre         string `json:"nombre"`
	Apellido       string `json:"apellido"`
	Cedula         string `json:"cedula"`
	InicialDoc     string `json:"inicial_doc"`
	Direccion      string `json:"direccion"`
	Telefono       string `json:"telefono"`
	TelfCasa       string `json:"telf_casa"`
	TelfAdic       string `json:"telf_adic"`
	StatusContrato string `json:"status_contrato"`
	Suscripcion    string `json:"suscripcion"`
}
