package models

import "testing"

func TestGrado(t *testing.T) {
	grado := &Grado{
		Nombre: "Grado en Sistemas",
	}
	tests := []struct {
		nombre   string
		obtenido string
		esperado string
	}{
		{
			nombre:   "Validar Nombre",
			obtenido: grado.Nombre,
			esperado: "Grado en Sistemas",
		},
	}
	for _, test := range tests {
		if test.obtenido != test.esperado {
			t.Errorf("%s: esperado %s, obtenido %s", test.nombre, test.esperado, test.obtenido)
		}
	}
}
