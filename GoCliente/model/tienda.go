package model

type Tienda struct {
	Codigo      string `json:"codigo,omitempty"`
	Nombre      string `json:"nombre,omitempty"`
	Telefono    int    `json:"telefono,omitempty"`
	Ciudad      string `json:"ciudad,omitempty"`
	Direccion   string `json:"direccion,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
}
