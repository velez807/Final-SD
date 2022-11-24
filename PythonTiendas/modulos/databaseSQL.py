from sqlalchemy import create_engine, MetaData, Table, Column, Integer, String

engine = create_engine('sqlite:///./PythonTiendas/modulos/tiendas.db', echo=True)

meta = MetaData()

conn = engine.connect()

tiendas = Table(
    'tiendas', meta,
    Column('codigo', Integer, primary_key=True),
    Column('nombre', String),
    Column('telefono', String),
    Column('ciudad', String),
    Column('direccion', String),
    Column('descripcion', String)
)