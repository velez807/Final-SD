from fastapi import Request, FastAPI
import sqlite3

# Base de datos SQLite3
con = sqlite3.connect("./PythonProductos/productos.db")

# ver la base de datos
con.row_factory = sqlite3.Row
cur = con.cursor()
cur.execute("SELECT * FROM productos")
rows = cur.fetchall()
for row in rows:
    print(dict(row))

# instanciamos aplicacion
#app = FastAPI()