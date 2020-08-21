package main

//DataWS es la estructura base para el retorno de información del api de precios de los almacenes
type DataWS struct {
	Almacen   string
	Productos []ProductoWS
}

//ProductoWS es la estructura base para el retorno de información del api de precios de los almacenes
type ProductoWS struct {
	Nombre string
	Precio string
	ImgURL string
}

//ResultDataPage es la estructura base para el envio de info a la pagina web de resultados
type ResultDataPage struct {
	Title      string
	SearchText string
	Data       []DataWS
}
