from peewee import *

# proxy
database_proxy = Proxy()

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
        database = database_proxy
        table_name = 'tiendas'

# base de datos
db = SqliteDatabase('./PythonTiendas/modulos/tiendas.db')

# conectar la base de datos
database_proxy.initialize(db)