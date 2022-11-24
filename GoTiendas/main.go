package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/velez807/Final-SD/handlers"
)

func main() {
	// instanciar el enrutador
	router := mux.NewRouter().StrictSlash(true)

	// rutas
	router.HandleFunc("/", handlers.ManejadorIndex).Methods("GET")
	router.HandleFunc("/tiendas", handlers.ManejadorConsultarTiendas).Methods("GET")
	router.HandleFunc("/tienda/{id}", handlers.ManejadorConsultarTiendaID).Methods("GET")
	router.HandleFunc("/tienda/{id}", handlers.ManejadorDeleteTiendaID).Methods("DELETE")
	router.HandleFunc("/tienda/", handlers.ManejadorPostTiendaID).Methods("POST")

	fmt.Println("Servidor iniciado en el puerto 8000")

	// iniciar el servidor
	log.Fatal(http.ListenAndServe(":8000", router))
}
