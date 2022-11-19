from pydantic import BaseModel


class ProductoRequestModel(BaseModel):
    codigo: str
    nombre: str
    precio: str
    descripcion: str
    tienda: str

class ProductoResponseModel(ProductoRequestModel):
    pass