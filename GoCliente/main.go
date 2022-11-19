package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/velez807/Final-SD/GoCliente/model"
)

//Acciones para obtener nota
func obtenerProductoID(idProd string, canalProductos chan *http.Response) {
	//Instanciar un cliente HTTP
	cliente := http.Client{}
	resp, err := cliente.Get("http://localhost:3000/consultarNota/" + idProd)
	if err != nil {
		log.Fatal("Error conmiendo servicio producto: ", idProd)
		os.Exit(1)
	}
	canalProductos <- resp
}

func obtenerTiendaID(idTienda string, canalTiendas chan *http.Response) {
	//Instanciar un cliente HTTP
	cliente := http.Client{}
	resp, err := cliente.Get("http://127.0.0.1:2323/estudiantesID?idTienda=" + idTienda)
	if err != nil {
		log.Fatal("Error conmiendo servicio tiendas: ", idTienda)
		os.Exit(1)
	}
	canalTiendas <- resp
}

func main() {
	var op int
	//Un canal por cada microservicio
	canalProductos := make(chan *http.Response)
	canalTiendas := make(chan *http.Response)

	var producto model.Producto
	var tienda model.Tienda
	fmt.Print("Menú de opciones: \n 1: Get por id \n Cosas \n Su opción: ")
	fmt.Scan(&op)
	switch op {
	case 1:
		//Consumos concurrentes
		var idtienda string
		var idprod string
		fmt.Print("Ingrese el ID del producto: ")
		fmt.Scan(&idprod)
		fmt.Print("Ingrese el ID de la tienda: ")
		fmt.Scan(&idtienda)
		go obtenerProductoID(idprod, canalProductos)
		go obtenerTiendaID(idtienda, canalTiendas)
		var datosProducto, datosTienda *http.Response
		//Vamos a realizar dos lecturas de los canales
		for i := 0; i < 2; i++ {
			select {
			case respProductos := <-canalProductos:
				datosProducto = respProductos
			case respTienda := <-canalTiendas:
				datosTienda = respTienda
			}

		}
		//Decodificar respuestas
		json.NewDecoder(datosProducto.Body).Decode(&producto)
		json.NewDecoder(datosTienda.Body).Decode(&tienda)
		//Interfaz
		fmt.Println("-------------------------------")
		fmt.Println("Producto consultado: ", producto)
		fmt.Println("Tienda consultada: ", tienda)
		fmt.Println("-------------------------------")
	}
}
