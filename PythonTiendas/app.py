from fastapi import FastAPI

# simular una base de datos de tiendas
bd = [
    {"codigo": 1, "nombre": "Juan", "telefono": "123456789", "ciudad": "Madrid", "direccion": "Calle 1", "descripcion": "Tienda 1"},
    {"codigo": 2, "nombre": "luis", "telefono": "123456789", "ciudad": "Madrid", "direccion": "Calle 2", "descripcion": "Tienda 2"},
    {"codigo": 3, "nombre": "julio", "telefono": "123456789", "ciudad": "Madrid", "direccion": "Calle 3", "descripcion": "Tienda 3"}
]

# instanciamos aplicacion
app = FastAPI()

@app.get("/tiendaCodigo")
def obtenerTiendaPorCodigo(codigo=None):
    
        if codigo == None:
            codigo = 1
    
        #controlar tipo
        codigo = int(codigo)
        
        #preparar tienda de retorno
        tiendaEncontrada = {}
    
        #buscar tienda
        for tienda in bd:
            if tienda["codigo"] == codigo:
                tiendaEncontrada = tienda
                break
    
        return tiendaEncontrada

