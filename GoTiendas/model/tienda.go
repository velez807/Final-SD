package model

import (
	"github.com/velez807/Final-SD/data"
)

type Tienda struct {
	Codigo      string `json:"codigo,omitempty"`
	Nombre      string `json:"nombre,omitempty"`
	Telefono    int    `json:"telefono,omitempty"`
	Ciudad      string `json:"ciudad,omitempty"`
	Direccion   string `json:"direccion,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
}

// metodo para obtener todas las tiendas
func (t Tienda) ObtenerTiendas() ([]Tienda, error) {
	// obtener la conexion
	db := data.ObtenerConexion("./data/tiendas.db")

	// query
	q := `SELECT * FROM tiendas`

	// realizar la consulta
	rows, err := db.Query(q)

	// revisar error
	if err != nil {
		return nil, err
	}

	// preparar la instancia de retorno
	var tiendas []Tienda

	// recorrer las filas
	for rows.Next() {
		// preparar la instancia
		var tienda Tienda

		// asignar valores
		err := rows.Scan(&tienda.Codigo, &tienda.Nombre, &tienda.Telefono, &tienda.Ciudad, &tienda.Direccion, &tienda.Descripcion)

		// revisar error
		if err != nil {
			return nil, err
		}

		// agregar a la lista
		tiendas = append(tiendas, tienda)
	}

	// retornar
	return tiendas, nil
}

// metodo para obtener tienda por id
func (t Tienda) ObtenerTiendaID(id string) (Tienda, error) {
	// obtener la conexion
	db := data.ObtenerConexion("./data/tiendas.db")

	// query
	q := `SELECT * FROM tiendas WHERE codigo = ?`

	// realizar la consulta
	row := db.QueryRow(q, id)

	// preparar la instancia de retorno
	var tiendas Tienda

	// asignar la respuesta a la estructura
	row.Scan(
		&tiendas.Codigo,
		&tiendas.Nombre,
		&tiendas.Telefono,
		&tiendas.Ciudad,
		&tiendas.Direccion,
		&tiendas.Descripcion,
	)

	// retornar
	return tiendas, nil
}
