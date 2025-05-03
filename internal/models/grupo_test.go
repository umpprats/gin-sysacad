package models

import "testing"

func TestGrupo(t *testing.T) {

	grupo := &Grupo{
		Nombre: "Grupo de Sistemas",
	}
	tests := []struct {
		nombre   string
		obtenido string
		esperado string
	}{
		{
			nombre:   "Validar Nombre",
			obtenido: grupo.Nombre,
			esperado: "Grupo de Sistemas",
		},
	}
	for _, test := range tests {
		if test.obtenido != test.esperado {
			t.Errorf("%s: esperado %s, obtenido %s", test.nombre, test.esperado, test.obtenido)
		}
	}
}
