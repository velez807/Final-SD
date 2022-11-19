from pydantic import BaseModel

class TiendaRequestModel(BaseModel):
    codigo: str
    nombre: str
    telefono: str
    ciudad: str
    direccion: str
    descripcion: str

class TiendaResponseModel(TiendaRequestModel):
    pass