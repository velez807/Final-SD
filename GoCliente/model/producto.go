package model

type Producto struct {
	codigo      int    `json:"codigo,omitempty"`
	nombre      string `json:"nombre,omitempty"`
	precio      int    `json:"precio,omitempty"`
	descripcion string `json:"descripcion,omitempty"`
	tienda      string `json:"tienda,omitempty"`
}
