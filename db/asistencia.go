package db

import "fmt"

type Asistencia struct {
	id            int
	nombre_alumno string
	tema_clase    string
}

func NewAsistencia(id int, nombre_alumno_ingresado string, nombre_tema string) Asistencia {
	return Asistencia{
		id:            id,
		nombre_alumno: nombre_alumno_ingresado,
		tema_clase:    nombre_tema,
	}
}
func InsertAsistencia(nuevaAsistencia Asistencia) (int, error) {
	//Create
	var asistenciaID int
	err := db.QueryRow(`INSERT INTO ASISTENCIA(NOMBREALUMNO,TEMACLASE) VALUES($1,$2) RETURNING IDASISTENCIA`, nuevaAsistencia.nombre_alumno, nuevaAsistencia.tema_clase).Scan(&asistenciaID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Last inserted ID: %v\n", asistenciaID)
	return asistenciaID, err
}

func GetAsistencias() ([]Asistencia, error) {
	asistencias := []Asistencia{}

	rows, err := db.Query(`SELECT * FROM ASISTENCIA;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id_actual int
		var nombre_alumno_actual string
		var tema_clase_actual string
		var AsistenciaActual Asistencia

		err = rows.Scan(&id_actual, &nombre_alumno_actual, &tema_clase_actual)
		if err != nil {
			return asistencias, err
		}
		AsistenciaActual = Asistencia{id: id_actual, nombre_alumno: nombre_alumno_actual, tema_clase: tema_clase_actual}
		asistencias = append(asistencias, AsistenciaActual)
	}
	return asistencias, err
}
func GetAsistenciasClase(nombre_clase string) ([]Asistencia, error) {
	asistencias := []Asistencia{}

	rows, err := db.Query(`SELECT * FROM ASISTENCIA WHERE TEMACLASE=$1;`, nombre_clase)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id_actual int
		var nombre_alumno_actual string
		var tema_clase_actual string
		var AsistenciaActual Asistencia

		err = rows.Scan(&id_actual, &nombre_alumno_actual, &tema_clase_actual)
		if err != nil {
			return asistencias, err
		}
		AsistenciaActual = Asistencia{id: id_actual, nombre_alumno: nombre_alumno_actual, tema_clase: tema_clase_actual}
		asistencias = append(asistencias, AsistenciaActual)
	}
	return asistencias, err
}
