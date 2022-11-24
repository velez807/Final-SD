package model

type Producto struct {
	Codigo      string    `json:"codigo,omitempty"`
	Nombre      string `json:"nombre,omitempty"`
	Precio      string    `json:"precio,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
	Tienda      string `json:"tienda,omitempty"`
}
