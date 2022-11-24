# autores:
# Andres Felipe Gutierrez
# Alejandro Buitrago

from fastapi import FastAPI, HTTPException
from modulos.database import Producto, database_proxy
from modulos.schemas import ProductoRequestModel, ProductoResponseModel
import uvicorn

# instanciamos aplicacion
app = FastAPI(title="API Productos", description="API para gestionar productos", version="1.0")

# endpoints como decoradores

@app.on_event("startup")
def startup():
    if database_proxy.is_closed():
        database_proxy.connect()
    

@app.on_event("shutdown")
def shutdown():
    if not database_proxy.is_closed():
        database_proxy.close()

@app.get("/")
async def index():
    return {"mensaje": "Bienvenido a la API de Productos"}

#crear
@app.post("/producto")
async def crear_producto(producto: ProductoRequestModel):
    producto = Producto.create(
        codigo=producto.codigo,
        nombre=producto.nombre,
        precio=producto.precio,
        descripcion=producto.descripcion,
        tienda=producto.tienda
    )
    if producto: 
        return ProductoResponseModel(codigo=producto.codigo, 
                                    nombre=producto.nombre, 
                                    precio=producto.precio, 
                                    descripcion=producto.descripcion, 
                                    tienda=producto.tienda)
    else:
        return HTTPException(200, 'Error en producto, revisar ID')

# actualizar
@app.put("/producto/{codigo}")
async def actualizar_producto(codigo: int, producto: ProductoRequestModel):
    producto = Producto.update(
        nombre=producto.nombre,
        precio=producto.precio,
        descripcion=producto.descripcion,
        tienda=producto.tienda
    ).where(Producto.codigo == codigo).execute()
    if producto:
        producto = Producto.select().where(Producto.codigo == codigo).first()
        return ProductoResponseModel(codigo=producto.codigo,
                                    nombre=producto.nombre,
                                    precio=producto.precio,
                                    descripcion=producto.descripcion,
                                    tienda=producto.tienda)
    else:
        return HTTPException(404, 'Producto no encontrado')


# obtener producto por codigo
@app.get("/producto/{codigo}")
async def obtener_producto(codigo: str):
    producto = Producto.select().where(Producto.codigo == codigo).first()
    if producto:
        return ProductoResponseModel(codigo=producto.codigo,
                                    nombre=producto.nombre,
                                    precio=producto.precio,
                                    descripcion=producto.descripcion,
                                    tienda=producto.tienda)
    else:
        return HTTPException(404, 'Producto no encontrado')

# obtener todos los productos
@app.get("/producto")
async def obtener_todos_los_productos():
    productos = Producto.select()
    return [ProductoResponseModel(codigo=producto.codigo, 
                                    nombre=producto.nombre,
                                    precio=producto.precio,
                                    descripcion=producto.descripcion,
                                    tienda=producto.tienda) for producto in productos]


#eliminar producto por codigo
@app.delete("/producto/{codigo}")
async def eliminar_producto(codigo: str):
    producto = Producto.select().where(Producto.codigo == codigo).first()
    if producto:
        producto.delete_instance()
        return {"mensaje": "producto eliminado con exito"}
    else:
        return HTTPException(404, 'Producti no encontrado en la BD')

# dejar esto siempre de ultimo
if __name__ == '__main__':
    uvicorn.run("API_Productos:app", port=3000, reload=True)