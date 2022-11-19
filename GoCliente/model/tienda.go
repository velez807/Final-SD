package model

type Tienda struct {
	codigo      string `json:"codigo,omitempty"`
	nombre      string `json:"nombre,omitempty"`
	telefono    int    `json:"telefono,omitempty"`
	ciudad      string `json:"ciudad,omitempty"`
	direccion   string `json:"direccion,omitempty"`
	descripcion string `json:"descripcion,omitempty"`
}
