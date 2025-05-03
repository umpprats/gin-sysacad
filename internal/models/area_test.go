package models

import "testing"

func TestArea(t *testing.T) {
	area := &Area{
		Nombre: "Area de Sistemas",
	}
	tests := []struct {
		nombre   string
		obtenido string
		esperado string
	}{
		{
			nombre:   "Validar Nombre",
			obtenido: area.Nombre,
			esperado: "Area de Sistemas",
		},
	}
	for _, test := range tests {
		if test.obtenido != test.esperado {
			t.Errorf("%s: esperado %s, obtenido %s", test.nombre, test.esperado, test.obtenido)
		}
	}
}
