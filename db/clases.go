package db

import (
	"fmt"
)

type Clase struct {
	id   int
	tema string
}

func NewClase(id_input int, tema_input string) Clase {
	return Clase{
		id:   id_input,
		tema: tema_input,
	}
}
func InsertClase(nuevoClase Clase) (int, error) {
	//Create
	var ClaseID int
	err := db.QueryRow(`INSERT INTO CLASE(TEMA) VALUES($1) RETURNING IDCLASE`, nuevoClase.tema).Scan(&ClaseID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Last inserted ID: %v\n", ClaseID)
	return ClaseID, err
}

func GetClases() ([]Clase, error) {
	Clases := []Clase{}

	rows, err := db.Query(`SELECT * FROM CLASE;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id_actual int
		var tema_actual string
		var ClaseActual Clase

		err = rows.Scan(&id_actual, &tema_actual)
		if err != nil {
			return Clases, err
		}
		ClaseActual = Clase{id: id_actual, tema: tema_actual}
		Clases = append(Clases, ClaseActual)
	}
	return Clases, err
}
