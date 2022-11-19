from fastapi import FastAPI
from modulos.database import Producto, db
from modulos.schemas import ProductoRequestModel, ProductoResponseModel
import uvicorn

# instanciamos aplicacion
app = FastAPI(title="API Productos", description="API para gestionar productos", version="1.0")

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
    return {"mensaje": "Bienvenido a la API de Productos"}



# dejar esto siempre de ultimo
if __name__ == '__main__':
    uvicorn.run(app, port=3000)