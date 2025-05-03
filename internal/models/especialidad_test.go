package models

import "testing"

func TestEspecialidad(t *testing.T) {
	especialidad := &Especialidad{
		Nombre:      "Tipo de Especialidad",
		Letra:       "A",
		Observacion: "Observacion de Especialidad",
		TipoEspecialidad: &TipoEspecialidad{
			Nombre: "Tipo de Especialidad",
			Nivel:  "I",
		},
	}
	tests := []struct {
		nombre   string
		obtenido string
		esperado string
	}{
		{
			nombre:   "Validar Nombre",
			obtenido: especialidad.Nombre,
			esperado: "Tipo de Especialidad",
		},
		{
			nombre:   "Validar Letra",
			obtenido: especialidad.Letra,
			esperado: "A",
		},
		{
			nombre:   "Validar Observacion",
			obtenido: especialidad.Observacion,
			esperado: "Observacion de Especialidad",
		},
		{
			nombre:   "Validar Nivel",
			obtenido: especialidad.TipoEspecialidad.Nivel,
			esperado: "I",
		},
	}
	for _, test := range tests {
		if test.obtenido != test.esperado {
			t.Errorf("%s: esperado %s, obtenido %s", test.nombre, test.esperado, test.obtenido)
		}
	}
}
