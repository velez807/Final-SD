from peewee import *

# base de datos
db = SqliteDatabase('productos.db')

# modelo de datos
class Producto(Model):
    codigo = IntegerField(primary_key=True)
    nombre = CharField()
    precio = IntegerField()
    descripcion = CharField()
    tienda = CharField()

    class Meta:
        database = db
        table_name = 'productos'