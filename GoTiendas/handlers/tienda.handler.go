package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/velez807/Final-SD/model"
)

// consultar todas las tiendas
func ManejadorConsultarTiendas(w http.ResponseWriter, r *http.Request) {
	// instanciar el modelo
	t := new(model.Tienda)

	// realizar la consulta
	tiendas, err := t.ObtenerTiendas()

	// revisar error
	if err != nil {
		log.Fatal("Error consultando la DB \n", err)
		os.Exit(1)
	}

	// revisar respuesta
	fmt.Println("respuesta: ", tiendas)

	// preparar la respuesta
	// convertir estructura a json
	respuesta, err := json.Marshal(tiendas)

	// revisar error
	if err != nil {
		log.Fatal("Error convirtiendo a json \n", err)
		os.Exit(1)
	}

	// enviar respuesta transmitiendo por http
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respuesta)

}

// manejador consultar tienda por id
func ManejadorConsultarTiendaID(w http.ResponseWriter, r *http.Request) {
	// instanciar el modelo
	t := new(model.Tienda)

	// obtener el id
	vars := mux.Vars(r)
	id := vars["id"]

	// revisar id recibido
	fmt.Println("id recibido: ", id)

	// realizar la consulta
	tienda, err := t.ObtenerTiendaID(id)

	// revisar error
	if err != nil {
		log.Fatal("Error consultando la DB \n", err)
		os.Exit(1)
	}

	// revisar respuesta
	fmt.Println("respuesta: ", tienda)

	// preparar la respuesta
	// convertir estructura a json
	respuesta, err := json.Marshal(tienda)

	// revisar error
	if err != nil {
		log.Fatal("Error convirtiendo a json \n", err)
		os.Exit(1)
	}

	// enviar respuesta transmitiendo por http
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respuesta)

}

// manejador delete tienda por id
func ManejadorDeleteTiendaID(w http.ResponseWriter, r *http.Request) {
	// instanciar el modelo
	t := new(model.Tienda)

	// obtener el id
	vars := mux.Vars(r)
	id := vars["id"]

	// revisar id recibido
	fmt.Println("id recibido: ", id)

	// realizar la consulta
	resp, err := t.DeleteTiendaID(id)

	// revisar error
	if err != nil {
		log.Fatal("Error consultando la DB \n", err)
		os.Exit(1)
	}

	// revisar respuesta
	fmt.Println("respuesta: ", resp)

	// enviar respuesta transmitiendo por http
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp))

}
// manejador delete tienda por id
func ManejadorPostTiendaID(w http.ResponseWriter, r *http.Request) {
	// instanciar el modelo
	t := new(model.Tienda)

	// obtener el json
	err := json.NewDecoder(r.Body).Decode(&t)
    if err != nil {
        log.Fatal(err)
		os.Exit(1)
    }
	

	// revisar id recibido
	fmt.Println("tienda recibida: ", *t)

	// realizar la consulta
	resp, err := t.PostTienda(*t)

	// revisar error
	if err != nil {
		log.Fatal("Error consultando la DB \n", err)
		os.Exit(1)
	}

	// revisar respuesta
	fmt.Println("respuesta: ", resp)

	// enviar respuesta transmitiendo por http
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp))

}

