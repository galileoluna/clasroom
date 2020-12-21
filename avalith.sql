drop database if exists AVALITH;

create database AVALITH;

CREATE TABLE IF NOT EXISTS ALUMNO(
IDALUMNO SERIAL,
NOMBRE varchar(30),
CONSTRAINT PK_ALUMNO PRIMARY KEY (IDALUMNO));

CREATE TABLE IF NOT EXISTS CLASE(
IDCLASE SERIAL,
TEMA varchar(30),
CONSTRAINT PK_CLASE PRIMARY KEY (IDCLASE));

CREATE TABLE IF NOT EXISTS ASISTENCIA(
IDASISTENCIA SERIAL,
NOMBREALUMNO varchar(30),
TEMACLASE varchar(30),
CONSTRAINT PK_ASISTENTE PRIMARY KEY (IDASISTENCIA));


INSERT INTO ALUMNO (NOMBRE) VALUES ('JOAO');
INSERT INTO ALUMNO (NOMBRE) VALUES ('NEHEMIAS');
INSERT INTO ALUMNO (NOMBRE) VALUES ('CELESTE');
INSERT INTO ALUMNO (NOMBRE) VALUES ('LUCIANA');


INSERT INTO CLASE (TEMA) VALUES ('CLASE-STRUCTS');
INSERT INTO CLASE (TEMA) VALUES ('CLASE-VARIABLES');
