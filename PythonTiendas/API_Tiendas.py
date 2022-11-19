# autores:
# Juan Sebastian Velez
# Javier Jurado


from fastapi import FastAPI, HTTPException
from modulos.database import Tienda, db
from modulos.schemas import TiendaRequestModel, TiendaResponseModel
import uvicorn

# instanciamos aplicacion
app = FastAPI(title="API Tiendas", description="API para gestionar tiendas", version="1.0")


# endpoints como decoradores
@app.on_event("startup")
def startup():
    if db.is_closed():
        db.connect()
    

@app.on_event("shutdown")
def shutdown():
    if not db.is_closed():
        db.close()

@app.get("/")
async def index():
    return {"mensaje": "Bienvenido a la API de Tiendas"}

#crear
@app.post("/tienda")
async def crear_tienda(tienda: TiendaRequestModel):
    tienda = Tienda.create(
        codigo=tienda.codigo,
        nombre=tienda.nombre,
        telefono=tienda.telefono,
        ciudad=tienda.ciudad,
        direccion=tienda.direccion,
        descripcion=tienda.descripcion
    )
    return tienda

# leer
@app.get("/tienda/{codigo}")
async def obtener_tienda(codigo: str):
    tienda = Tienda.select().where(Tienda.codigo == codigo).first()
    if tienda:
        return TiendaResponseModel(codigo=tienda.codigo, 
                                    nombre=tienda.nombre, 
                                    telefono=tienda.telefono, 
                                    ciudad=tienda.ciudad, 
                                    direccion=tienda.direccion, 
                                    descripcion=tienda.descripcion)
    else:
        return HTTPException(404, 'Tienda no encontrada')

# obtener todas las tiendas
@app.get("/tiendas")
async def obtener_todas_las_tiendas():
    tiendas = Tienda.select()
    return [TiendaResponseModel(codigo=tienda.codigo, 
                                nombre=tienda.nombre, 
                                telefono=tienda.telefono, 
                                ciudad=tienda.ciudad, 
                                direccion=tienda.direccion, 
                                descripcion=tienda.descripcion) for tienda in tiendas]


# actualizar
@app.put("/tienda/{codigo}")
async def actualizar_tienda(codigo: str, tienda: TiendaRequestModel):
    tienda = Tienda.update(
        nombre=tienda.nombre,
        telefono=tienda.telefono,
        ciudad=tienda.ciudad,
        direccion=tienda.direccion,
        descripcion=tienda.descripcion
    ).where(Tienda.codigo == codigo).execute()
    if tienda:
        tienda = Tienda.select().where(Tienda.codigo == codigo).first()
        return TiendaResponseModel(codigo=tienda.codigo, 
                                    nombre=tienda.nombre, 
                                    telefono=tienda.telefono, 
                                    ciudad=tienda.ciudad, 
                                    direccion=tienda.direccion, 
                                    descripcion=tienda.descripcion)
    else:
        return HTTPException(404, 'Tienda no encontrada')

#eliminar
@app.delete("/tienda/{codigo}")
async def eliminar_tienda(codigo: str):
    tienda = Tienda.select().where(Tienda.codigo == codigo).first()
    if tienda:
        tienda.delete_instance()
        return {"mensaje": "Tienda eliminada"}
    else:
        return HTTPException(404, 'Tienda no encontrada')


# dejar esto siempre de ultimo
if __name__ == '__main__':
    uvicorn.run("API_Tiendas:app", port=8000, reload=True)