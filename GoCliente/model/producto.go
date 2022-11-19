package model

type Producto struct {
	Codigo      int    `json:"codigo,omitempty"`
	Nombre      string `json:"nombre,omitempty"`
	Precio      int    `json:"precio,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
	Tienda      string `json:"tienda,omitempty"`
}
