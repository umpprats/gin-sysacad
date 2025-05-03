package models

import "testing"

func TestMateria(t *testing.T) {
	// Create a new Materia instance
	materia := &Materia{
		Nombre:      "Programacion I",
		Codigo:      "101",
		Observacion: "Obligatoria",
	}

	tests := []struct {
		nombre   string
		obtenido string
		esperado string
	}{
		{
			nombre:   "Validar Nombre",
			obtenido: materia.Nombre,
			esperado: "Programacion I",
		},
		{
			nombre:   "Validar Codigo",
			obtenido: materia.Codigo,
			esperado: "101",
		},
		{
			nombre:   "Validar Observacion",
			obtenido: materia.Observacion,
			esperado: "Obligatoria",
		},
	}

	for _, test := range tests {
		if test.obtenido != test.esperado {
			t.Errorf("%s: esperado %s, obtenido %s", test.nombre, test.esperado, test.obtenido)
		}
	}

}
