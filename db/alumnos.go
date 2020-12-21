package db

import (
	"fmt"
)

type Alumno struct {
	id     int
	nombre string
}

func NewAlumno(id_input int, nombre_input string) Alumno {
	return Alumno{
		id:     id_input,
		nombre: nombre_input,
	}
}

func InsertAlumno(nuevoAlumno Alumno) (int, error) {
	//Create
	var alumnoID int
	err := db.QueryRow(`INSERT INTO ALUMNO(NOMBRE) VALUES($1) RETURNING IDALUMNO`, nuevoAlumno.nombre).Scan(&alumnoID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Last inserted ID: %v\n", alumnoID)
	return alumnoID, err
}

func GetAlumnos() ([]Alumno, error) {
	alumnos := []Alumno{}
	rows, err := db.Query(`SELECT IDALUMNO,NOMBRE FROM ALUMNO`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var IDALUMNO int
		var NOMBRE string
		var alumnoActual Alumno

		err = rows.Scan(&IDALUMNO, &NOMBRE)
		if err != nil {
			return alumnos, err
		}
		alumnoActual = Alumno{id: IDALUMNO, nombre: NOMBRE}
		alumnos = append(alumnos, alumnoActual)
	}
	return alumnos, err
}
