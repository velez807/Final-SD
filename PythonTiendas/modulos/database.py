from peewee import *

# base de datos
db = SqliteDatabase('./PythonTiendas/modulos/tiendas.db')

class Tienda(Model):
    codigo = CharField(primary_key=True)
    nombre = CharField()
    telefono = CharField()
    ciudad = CharField()
    direccion = CharField()
    descripcion = CharField()

    def __str__(self):
        return self.nombre

    class Meta:
        database = db
        table_name = 'tiendas'