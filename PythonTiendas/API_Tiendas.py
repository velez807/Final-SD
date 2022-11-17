from fastapi import FastAPI, HTTPException
from database import Tienda
from database import db
from schemas import TiendaRequestModel, TiendaResponseModel

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

# actualizar
@app.put("/tienda/{codigo}")
async def actualizar_tienda(codigo: str, tienda: TiendaRequestModel):
    tienda = Tienda.update(
        codigo=tienda.codigo,
        nombre=tienda.nombre,
        telefono=tienda.telefono,
        ciudad=tienda.ciudad,
        direccion=tienda.direccion,
        descripcion=tienda.descripcion
    ).where(Tienda.codigo == codigo)
    return HTTPException(200, 'Tienda actualizada')




#eliminar
@app.delete("/tienda/{codigo}")
async def obtener_tienda(codigo: str):
    tienda = Tienda.select().where(Tienda.codigo == codigo).first()
    if tienda:
        tienda.delete_instance()
        return {"mensaje": "Tienda eliminada"}
    else:
        return HTTPException(404, 'Tienda no encontrada')