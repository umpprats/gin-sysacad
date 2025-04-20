package models

import (
	"strconv"
	"testing"
	"time"
)

func TestAlumno(t *testing.T) {
	// Test cases for the Alumno model
	alumno := &Alumno{
		Apellido:        "Prats",
		Nombre:          "Pablo Andres",
		NroDocumento:    "12345678",
		TipoDocumento:   "DNI",
		FechaNacimiento: time.Date(1975, 1, 1, 0, 0, 0, 0, time.UTC),
		Sexo:            "M",
		NroLegajo:       123456,
		FechaIngreso:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	tests := []struct {
		nombre   string
		obtenido string
		esperado string
	}{
		{
			nombre:   "Validar Nombre",
			obtenido: alumno.Nombre,
			esperado: "Pablo Andres",
		},
		{
			nombre:   "Validar Apellido",
			obtenido: alumno.Apellido,
			esperado: "Prats",
		},
		{
			nombre:   "Validar NroDocumento",
			obtenido: alumno.NroDocumento,
			esperado: "12345678",
		},
		{
			nombre:   "Validar TipoDocumento",
			obtenido: string(alumno.TipoDocumento),
			esperado: "DNI",
		},
		{
			nombre:   "Validar FechaNacimiento",
			obtenido: alumno.FechaNacimiento.Format("2006-01-02"),
			esperado: "1975-01-01",
		},
		{
			nombre:   "Validar Sexo",
			obtenido: alumno.Sexo,
			esperado: "M",
		},
		{
			nombre:   "Validar NroLegajo",
			obtenido: strconv.Itoa(alumno.NroLegajo),
			esperado: "123456",
		},
		{
			nombre:   "Validar FechaIngreso",
			obtenido: alumno.FechaIngreso.Format("2006-01-02"),
			esperado: "2023-01-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.nombre, func(t *testing.T) {
			if tt.obtenido != tt.esperado {
				t.Errorf("Se esperaba que %s tenga '%s', pero tiene '%s'", tt.nombre, tt.esperado, tt.obtenido)
			}
		})
	}

	// Validar que un alumno tiene el tipo de documento que pertenece a TipoDocumento utlizando el método IsValid()
	t.Run("Validar Tipo de Documento se un Tipo Documento Valido", func(t *testing.T) {
		if !alumno.TipoDocumento.IsValid() {
			t.Errorf("Se esperaba que el tipo de documento '%s' sea válido", alumno.TipoDocumento)
		}
	},
	)

}
