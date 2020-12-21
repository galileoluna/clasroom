package db_test

import (
	"testing"

	"github.com/galileoluna/clasroom/db"
)

func TestGetClasesSize(t *testing.T) {
	var clases []db.Clase

	var largoClases int

	clases, _ = db.GetClases()

	largoClases = len(clases)

	if largoClases == 0 {
		t.Error("No existen clases....")
	}

	if largoClases > 0 {
		t.Error("Existen clases")
	}
}
