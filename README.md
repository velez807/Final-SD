# Final-SD
<p>Proyecto final Sistemas distribuidos: microservicios y api rest en python y go. <br>
Universidad libre seccional Pereira.<br>
Ingeniería de Sistemas.<br>
2022</p>

# Autores:
- Juan Sebastián Vélez
- Julio Alejandro Peñaloza
- Javier jurado 
- Andres Felipe Gutierrez
- Alejandro Buitrago

# Instrucciones:
Instalar todas las librerías del archivo requirements.txt

# Aviso:
PythonTiendas funciona (se puede comprobar desde localhost:8000/docs) pero por alguna razón uvicorn no permite dos conexiones a base de datos al mismo tiempo, por lo cual no se podían ejecutar ambos archivos python juntos, en su defecto se implementó la api de tiendas en Go con gorilla/mux y net/http. <br>
Intentando solucionar el error se intentó cambiando de orm para usar dos diferentes (Peewee y SQLAlchemy) pero el error persistía, a pesar de que ya no se utiliza esta API igual decidimos dejarla para su evaluacion ya que fue trabajo de Javier y Sebastián

# Ejecucion
Ejecutar desde terminal los archivos Python API_Productos.py de PythonProductos (normal, sin uvicorn) y API_Tienda.go de GoTiendas <br>
API_Productos se aloja en localhost:3000.<br>
API_Tienda se aloja en localhost:8000.<br>
Las consultas y solicitudes concurrentes se encuentran en GoCliente. Pero las apis se pueden comprobar detenidamente desde Thunderclient <br>
