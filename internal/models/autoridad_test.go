package models

import "testing"

func TestAutoridad(t *testing.T) {

	autoridad := &Autoridad{
		Nombre:   "Ing. Roberto D. Vilches",
		Cargo:    &Cargo{Nombre: "Decano", Puntos: 10},
		Telefono: "02604421078",
		Email:    "rvilches@frsr.utn.edu.ar",
	}
	tests := []struct {
		nombre   string
		obtenido string
		esperado string
	}{
		{
			nombre:   "Validar Nombre",
			obtenido: autoridad.Nombre,
			esperado: "Ing. Roberto D. Vilches",
		},
		{
			nombre:   "Validar Cargo",
			obtenido: autoridad.Cargo.Nombre,
			esperado: "Decano",
		},
		{
			nombre:   "Validar Telefono",
			obtenido: autoridad.Telefono,
			esperado: "02604421078",
		},
		{
			nombre:   "Validar Email",
			obtenido: autoridad.Email,
			esperado: "rvilches@frsr.utn.edu.ar",
		},
	}
	for _, test := range tests {
		if test.obtenido != test.esperado {
			t.Errorf("%s: esperado %s, obtenido %s", test.nombre, test.esperado, test.obtenido)
		}
	}
}
