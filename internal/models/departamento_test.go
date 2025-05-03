package models

import "testing"

func TestDepartamento(t *testing.T) {

	// Create a new Departamento instance
	departamento := &Departamento{
		Nombre: "Departamento de Sistemas",
	}

	tests := []struct {
		nombre   string
		obtenido string
		esperado string
	}{
		{
			nombre:   "Validar Nombre",
			obtenido: departamento.Nombre,
			esperado: "Departamento de Sistemas",
		},
	}

	for _, test := range tests {
		if test.obtenido != test.esperado {
			t.Errorf("%s: esperado %s, obtenido %s", test.nombre, test.esperado, test.obtenido)
		}
	}
}
