package handlers

import (
	"fmt"
	"net/http"
)

func ManejadorIndex(w http.ResponseWriter, r *http.Request) {
	// respuesta por consola donde está corriendo el servicio
	fmt.Println("Index del servicio alcanzado")
	// respuesta a traves de http
	fmt.Fprintf(w, "Bienvenido al servicio de la clase index desde golang")
}
