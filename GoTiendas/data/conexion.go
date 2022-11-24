package data

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func ObtenerConexion(rutaDB string) *sql.DB {

	if db != nil {
		return db
	}

	var err error

	db, err = sql.Open("sqlite3", rutaDB)

	if err != nil {
		fmt.Println("error de conexion")
		panic(err)

	}

	return db
}
