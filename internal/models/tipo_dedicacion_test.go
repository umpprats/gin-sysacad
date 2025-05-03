package models

import "testing"

func TestTipoDedicacion(t *testing.T) {
	// Create a new TipoDedicacion instance
	tipoDedicacion := &TipoDedicacion{
		Nombre:      "Tiempo Completo",
		Observacion: "Dedicacion completa",
	}

	tests := []struct {
		nombre   string
		obtenido string
		esperado string
	}{
		{
			nombre:   "Validar Nombre",
			obtenido: tipoDedicacion.Nombre,
			esperado: "Tiempo Completo",
		},
		{
			nombre:   "Validar Observacion",
			obtenido: tipoDedicacion.Observacion,
			esperado: "Dedicacion completa",
		},
	}

	for _, test := range tests {
		if test.obtenido != test.esperado {
			t.Errorf("%s: esperado %s, obtenido %s", test.nombre, test.esperado, test.obtenido)
		}
	}
}
