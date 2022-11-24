//Autor Julio Peñaloza :D
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/velez807/Final-SD/GoCliente/model"
)

// Acciones para obtener nota
func obtenerProductoID(idProd string, canalProductos chan *http.Response) {
	//Instanciar un cliente HTTP
	cliente := http.Client{}
	resp, err := cliente.Get("http://127.0.0.1:3000/producto/" + idProd)
	if err != nil {
		log.Fatal("Error servicio producto: ", idProd)
		os.Exit(1)
	}
	canalProductos <- resp
}

func obtenerTiendaID(idTienda string, canalTiendas chan *http.Response) {
	//Instanciar un cliente HTTP
	cliente := http.Client{}
	resp, err := cliente.Get("http://127.0.0.1:8000/tienda/" + idTienda)
	if err != nil {
		log.Fatal("Error conmiendo servicio tiendas: ", idTienda)
		os.Exit(1)
	}
	canalTiendas <- resp
}
func PostTienda(tienda model.Tienda, canalTiendas chan *http.Response) {
	//Instanciar un cliente HTTP
	cliente := http.Client{}
	jsonTienda, err := json.Marshal(tienda)
	if err != nil {
		// Maneja el error de acuerdo a tu situación
		log.Fatalf("Error codificando usuario como JSON: %v", err)
	}
	url := "http://127.0.0.1:8000/tienda/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonTienda))
	if err != nil {
		log.Fatal("Error conmiendo servicio tiendas: ")
		os.Exit(1)
	}
	resp, err := cliente.Do(req)
	if err != nil {
		log.Fatal("Error conmiendo servicio tiendas: ")
		os.Exit(1)
	}
	canalTiendas <- resp
}
func PostProducto(producto model.Producto, canalProductos chan *http.Response) {
	//Instanciar un cliente HTTP
	cliente := http.Client{}
	jsonProducto, err := json.Marshal(producto)
	if err != nil {
		// Maneja el error de acuerdo a tu situación
		log.Fatalf("Error codificando usuario como JSON: %v", err)
	}
	url := "http://127.0.0.1:3000/producto/"
	fmt.Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonProducto))
	if err != nil {
		log.Fatal("Error arriba: ")
		os.Exit(1)
	}
	resp, err := cliente.Do(req)
	if err != nil {
		log.Fatal("Error servicio productos: ", err.Error())
		os.Exit(1)
	}
	canalProductos <- resp
}
func deleteTiendaID(idTienda string, canalRespuestas chan *http.Response) {
	cliente := http.Client{}
	url := "http://127.0.0.1:8000/tienda/" + idTienda
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal("Error conmiendo servicio tiendas: ", idTienda)
		os.Exit(1)
	}
	resp, err := cliente.Do(req)
	if err != nil {
		log.Fatal("Error conmiendo servicio tiendas: ", idTienda)
		os.Exit(1)
	}
	canalRespuestas <- resp
}
func deleteProductoID(idProd string, canalRespuestas chan *http.Response) {
	cliente := http.Client{}
	url := "http://127.0.0.1:3000/producto/" + idProd
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal("Error servicio productos: ", idProd)
		os.Exit(1)
	}
	resp, err := cliente.Do(req)
	if err != nil {
		log.Fatal("Error servicio productos: ", idProd)
		os.Exit(1)
	}
	canalRespuestas <- resp
}
func PutProducto(producto model.Producto, canalProductos chan *http.Response) {
	cliente := http.Client{}
	jsonProducto, err := json.Marshal(producto)
	if err != nil {
		// Maneja el error de acuerdo a tu situación
		log.Fatalf("Error codificando usuario como JSON: %v", err)
	}
	url := "http://127.0.0.1:3000/producto/" + strconv.Itoa(producto.Codigo)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonProducto))
	if err != nil {
		log.Fatal("Error arriba: ")
		os.Exit(1)
	}
	resp, err := cliente.Do(req)
	if err != nil {
		log.Fatal("Error servicio productos: ", err.Error())
		os.Exit(1)
	}
	canalProductos <- resp

}
func PutTienda(tienda model.Tienda, canalTiendas chan *http.Response) {
	cliente := http.Client{}
	jsonTienda, err := json.Marshal(tienda)
	if err != nil {
		// Maneja el error de acuerdo a tu situación
		log.Fatalf("Error codificando usuario como JSON: %v", err)
	}
	url := "http://127.0.0.1:8000/tienda/" + tienda.Codigo
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonTienda))
	if err != nil {
		log.Fatal("Error conmiendo servicio tiendas: ")
		os.Exit(1)
	}
	resp, err := cliente.Do(req)
	if err != nil {
		log.Fatal("Error conmiendo servicio tiendas: ")
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
	for op != 5 {
		fmt.Print("Menú de opciones: \n 1: Get por id \n 2: Delete por id \n 3: Post \n 4: Put \n 5: Salir \n Su opción: ")
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
		case 2:
			var idtienda string
			var idprod string
			fmt.Print("Ingrese el ID del producto: ")
			fmt.Scan(&idprod)
			fmt.Print("Ingrese el ID de la tienda: ")
			fmt.Scan(&idtienda)
			var resProducto, resTienda *http.Response
			go deleteTiendaID(idtienda, canalTiendas)
			go deleteProductoID(idprod, canalProductos)
			for i := 0; i < 2; i++ {
				select {
				case respProductos := <-canalProductos:
					resProducto = respProductos
				case respTienda := <-canalTiendas:
					resTienda = respTienda
				}
			}
			// Read Response Body
			respBody, err := ioutil.ReadAll(resTienda.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Display Results
			fmt.Println("response Status : ", resTienda.Status)
			fmt.Println("response Headers : ", resTienda.Header)
			fmt.Println("response Body : ", string(respBody))
			respBody2, err := ioutil.ReadAll(resProducto.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Display Results
			fmt.Println("response Status : ", resProducto.Status)
			fmt.Println("response Headers : ", resProducto.Header)
			fmt.Println("response Body : ", string(respBody2))
		case 3:
			var codigoTienda string
			var nombreTienda string
			var telefonoTienda int
			var ciudadTienda string
			var direccionTienda string
			var descripcionTienda string
			fmt.Print("Ingrese el ID de la tienda: ")
			fmt.Scan(&codigoTienda)
			fmt.Print("Ingrese el nombre de la tienda: ")
			fmt.Scan(&nombreTienda)
			fmt.Print("Ingrese el telefono de la tienda: ")
			fmt.Scan(&telefonoTienda)
			fmt.Print("Ingrese la ciudad de la tienda: ")
			fmt.Scan(&ciudadTienda)
			fmt.Print("Ingrese la direccion de la tienda: ")
			fmt.Scan(&direccionTienda)
			fmt.Print("Ingrese la descripcion de la tienda: ")
			fmt.Scan(&descripcionTienda)
			tiendaStr := model.Tienda{Codigo: codigoTienda, Nombre: nombreTienda, Telefono: telefonoTienda, Ciudad: ciudadTienda, Direccion: direccionTienda, Descripcion: descripcionTienda}
			fmt.Println(tiendaStr)
			var codigoProducto int
			var nombreProducto string
			var precioProducto int
			var descripcionProducto string
			var tiendaProducto string
			fmt.Print("Ingrese el ID del producto: ")
			fmt.Scan(&codigoProducto)
			fmt.Print("Ingrese el nombre del producto: ")
			fmt.Scan(&nombreProducto)
			fmt.Print("Ingrese el precio del producto: ")
			fmt.Scan(&precioProducto)
			fmt.Print("Ingrese la descripcion del producto: ")
			fmt.Scan(&descripcionProducto)
			fmt.Print("Ingrese la tienda del producto: ")
			fmt.Scan(&tiendaProducto)
			prodStr := model.Producto{Codigo: codigoProducto, Nombre: nombreProducto, Precio: precioProducto, Descripcion: descripcionProducto, Tienda: tiendaProducto}
			var resProducto, resTienda *http.Response
			go PostTienda(tiendaStr, canalTiendas)
			go PostProducto(prodStr, canalProductos)
			for i := 0; i < 2; i++ {
				select {
				case respProductos := <-canalProductos:
					resProducto = respProductos
				case respTienda := <-canalTiendas:
					resTienda = respTienda
				}
			}
			// Read Response Body
			respBody, err := ioutil.ReadAll(resTienda.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Display Results
			fmt.Println("response Status : ", resTienda.Status)
			fmt.Println("response Headers : ", resTienda.Header)
			fmt.Println("response Body : ", string(respBody))
			respBody2, err := ioutil.ReadAll(resProducto.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Display Results
			fmt.Println("response Status : ", resProducto.Status)
			fmt.Println("response Headers : ", resProducto.Header)
			fmt.Println("response Body : ", string(respBody2))
		case 4:
			var codigoTienda string
			var nombreTienda string
			var telefonoTienda int
			var ciudadTienda string
			var direccionTienda string
			var descripcionTienda string
			fmt.Print("Ingrese el ID de la tienda: ")
			fmt.Scan(&codigoTienda)
			fmt.Print("Ingrese el nombre de la tienda: ")
			fmt.Scan(&nombreTienda)
			fmt.Print("Ingrese el telefono de la tienda: ")
			fmt.Scan(&telefonoTienda)
			fmt.Print("Ingrese la ciudad de la tienda: ")
			fmt.Scan(&ciudadTienda)
			fmt.Print("Ingrese la direccion de la tienda: ")
			fmt.Scan(&direccionTienda)
			fmt.Print("Ingrese la descripcion de la tienda: ")
			fmt.Scan(&descripcionTienda)
			tiendaStr := model.Tienda{Codigo: codigoTienda, Nombre: nombreTienda, Telefono: telefonoTienda, Ciudad: ciudadTienda, Direccion: direccionTienda, Descripcion: descripcionTienda}
			fmt.Println(tiendaStr)
			var codigoProducto int
			var nombreProducto string
			var precioProducto int
			var descripcionProducto string
			var tiendaProducto string
			fmt.Print("Ingrese el ID del producto: ")
			fmt.Scan(&codigoProducto)
			fmt.Print("Ingrese el nombre del producto: ")
			fmt.Scan(&nombreProducto)
			fmt.Print("Ingrese el precio del producto: ")
			fmt.Scan(&precioProducto)
			fmt.Print("Ingrese la descripcion del producto: ")
			fmt.Scan(&descripcionProducto)
			fmt.Print("Ingrese la tienda del producto: ")
			fmt.Scan(&tiendaProducto)
			prodStr := model.Producto{Codigo: codigoProducto, Nombre: nombreProducto, Precio: precioProducto, Descripcion: descripcionProducto, Tienda: tiendaProducto}
			var resProducto, resTienda *http.Response
			go PutTienda(tiendaStr, canalTiendas)
			go PutProducto(prodStr, canalProductos)
			for i := 0; i < 2; i++ {
				select {
				case respProductos := <-canalProductos:
					resProducto = respProductos
				case respTienda := <-canalTiendas:
					resTienda = respTienda
				}
			}
			// Read Response Body
			respBody, err := ioutil.ReadAll(resTienda.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Display Results
			fmt.Println("response Status : ", resTienda.Status)
			fmt.Println("response Headers : ", resTienda.Header)
			fmt.Println("response Body : ", string(respBody))
			respBody2, err := ioutil.ReadAll(resProducto.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Display Results
			fmt.Println("response Status : ", resProducto.Status)
			fmt.Println("response Headers : ", resProducto.Header)
			fmt.Println("response Body : ", string(respBody2))
		case 5:
			fmt.Println("Adiosito")
		default:
			fmt.Println("Opción inválida")
		}
	}
}
