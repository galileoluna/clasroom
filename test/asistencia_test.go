package db_test

import (
	"testing"

	"github.com/galileoluna/clasroom/db"
)

func TestGetAsistenciaSize(t *testing.T) {
	var Asistencias []db.Asistencia

	var largoAsistencias int

	Asistencias, _ = db.GetAsistencias()

	largoAsistencias = len(Asistencias)

	if largoAsistencias == 0 {
		t.Error("No existen asistencias....")
	}

	if largoAsistencias > 0 {
		t.Error("Existen asistencias")
	}
}
