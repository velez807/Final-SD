from peewee import *

# proxy
database_proxy = Proxy()

# modelo de datos
class Producto(Model):
    codigo = IntegerField(primary_key=True)
    nombre = CharField()
    precio = IntegerField()
    descripcion = CharField()
    tienda = CharField()

    class Meta:
        database = database_proxy
        table_name = 'productos'

# base de datos
db = SqliteDatabase('./PythonProductos/modulos/productos.db')

# conectar la base de datos
database_proxy.initialize(db)
