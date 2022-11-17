-- tabla de tiendas con codigo, nombre, telefono, ciudad y direccion
CREATE TABLE tiendas (
  codigo VARCHAR PRIMARY KEY,
  nombre VARCHAR(50),
  telefono VARCHAR(20),
  ciudad VARCHAR(50),
  direccion VARCHAR(50),
  descripcion VARCHAR(150)
); 

-- insertar 5 tiendas diferentes en la tabla de tiendas
INSERT INTO tiendas (codigo, nombre, telefono, ciudad, direccion, descripcion) VALUES ('ELEC', 'Electro', '123456789', 'Pereira', 'Arboleda', 'Tienda de electronica');
INSERT INTO tiendas (codigo, nombre, telefono, ciudad, direccion, descripcion) VALUES ('ROPA', 'Ropa', '321459265', 'Dosquebradas', 'plaza bonita', 'Tienda de ropa');
INSERT INTO tiendas (codigo, nombre, telefono, ciudad, direccion, descripcion) VALUES ('ALIM', 'Alimentos', '987654321', 'Cartago', 'CC nuestro cartago', 'Tienda de alimentos');
INSERT INTO tiendas (codigo, nombre, telefono, ciudad, direccion, descripcion) VALUES ('JUGU', 'Juguetes', '123987456', 'La Virginia', 'Centro', 'Tienda de juguetes');
INSERT INTO tiendas (codigo, nombre, telefono, ciudad, direccion, descripcion) VALUES ('LIBR', 'Libreria', '456789123', 'Pereira', 'Unicentro', 'Tienda de libros');


-- tabla de productos con codigo, nombre, precio, descripcion, codigo de tienda
CREATE TABLE productos (
  codigo INT PRIMARY KEY,
  nombre VARCHAR(50),
  precio INT NOT NULL,
  descripcion VARCHAR(150),
  tienda VARCHAR(4)
);

-- 5 productos por cada tienda
-- tienda de electronica:
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (1, 'Laptop', 1000000, 'Laptop HP', 'ELEC');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (2, 'Celular', 500000, 'Celular Samsung', 'ELEC');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (3, 'Tablet', 300000, 'Tablet Samsung', 'ELEC');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (4, 'Televisor', 2000000, 'Television LG', 'ELEC');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (5, 'Audifonos', 100000, 'Audifonos Sony', 'ELEC');

-- tienda de ropa:
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (6, 'Camisa', 50000, 'Camisa de gucci', 'ROPA');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (7, 'Pantalon', 100000, 'Pantalon de luis vuiton', 'ROPA');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (8, 'Zapatos', 200000, 'Zapatos de giorgio armani', 'ROPA');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (9, 'Camiseta', 30000, 'Camiseta de puma', 'ROPA');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (10, 'Polo', 40000, 'Polo de Polo', 'ROPA');

-- tienda de alimentos:
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (11, 'Arroz', 10000, '2 kilos de arroz', 'ALIM');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (12, 'Frijoles', 20000, '3 kios de frijoles', 'ALIM');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (13, 'Leche', 3000, '1 caja de leche', 'ALIM');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (14, 'Huevos', 15000, '1 panal de huevos', 'ALIM');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (15, 'Carne', 14000, '1 libra de carne de res', 'ALIM');

-- tienda de juguetes:
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (16, 'Muñeca', 50000, 'Muñeca de barbie', 'JUGU');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (17, 'Pelota', 10000, 'Pelota de futbol', 'JUGU');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (18, 'Monopoly', 20000, 'Juego de mesa de monopoly', 'JUGU');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (19, 'Parques', 30000, 'Juego de mesa de parques', 'JUGU');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (20, 'Ajedrez', 40000, 'Juego de mesa de ajedrez', 'JUGU');

-- tienda de libros:
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (21, 'Harry Potter', 50000, 'Libro de harry potter', 'LIBR');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (22, 'El principito', 10000, 'Libro de el principito', 'LIBR');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (23, 'El alquimista', 20000, 'Libro de el alquimista', 'LIBR');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (24, 'El diario de Ana Frank', 30000, 'Libro de el diario de ana frank', 'LIBR');
INSERT INTO productos (codigo, nombre, precio, descripcion, tienda) VALUES (25, 'El señor de los anillos', 40000, 'Libro de el señor de los anillos', 'LIBR');