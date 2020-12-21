package db_test

import (
	"testing"

	"github.com/galileoluna/clasroom/db"
)

func TestGetAlumnoSize(t *testing.T) {
	var alumnos []db.Alumno

	var largoAlumnos int

	alumnos, _ = db.GetAlumnos()

	largoAlumnos = len(alumnos)

	if largoAlumnos == 0 {
		t.Error("No existentn alumnos....")
	}

	if largoAlumnos > 0 {
		t.Error("Existen alumnos")
	}
}
