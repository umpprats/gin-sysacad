package models

import "testing"

func TestFacultad(t *testing.T) {

	facultad := &Facultad{
		Nombre:       "Facultad Regional San Rafael",
		Abreviatura:  "FRSR",
		Directorio:   "Facultad Regional San Rafael",
		Sigla:        "FRSR",
		CodigoPostal: "5600",
		Ciudad:       "San Rafael",
		Domicilio:    "Av. Gral. J. J. Urquiza 314",
		Telefono:     "02604421078",
		Contacto:     "Ing. Roberto D. Vilches",
		Email:        "rvilches@frsr.utn.edu.ar",
	}

	t.Run("Test Nombre", func(t *testing.T) {
		if facultad.Nombre != "Facultad Regional San Rafael" {
			t.Errorf("Se esperaba que Nombre tenga 'Facultad Regional San Rafael', pero tiene '%s'", facultad.Nombre)
		}
	})

	t.Run("Test Abreviatura", func(t *testing.T) {
		if facultad.Abreviatura != "FRSR" {
			t.Errorf("Se esperaba que Abreviatura tenga 'FRSR', pero tiene '%s'", facultad.Abreviatura)
		}
	})

	t.Run("Test Directorio", func(t *testing.T) {
		if facultad.Directorio != "Facultad Regional San Rafael" {
			t.Errorf("Se esperaba que Directorio tenga 'Facultad Regional San Rafael', pero tiene '%s'", facultad.Directorio)
		}
	})

	t.Run("Test Sigla", func(t *testing.T) {
		if facultad.Sigla != "FRSR" {
			t.Errorf("Se esperaba que Sigla tenga 'FRSR', pero tiene '%s'", facultad.Sigla)
		}
	})

	t.Run("Test CodigoPostal", func(t *testing.T) {
		if facultad.CodigoPostal != "5600" {
			t.Errorf("Se esperaba que CodigoPostal tenga '5600', pero tiene '%s'", facultad.CodigoPostal)
		}
	})

	t.Run("Test Ciudad", func(t *testing.T) {
		if facultad.Ciudad != "San Rafael" {
			t.Errorf("Se esperaba que Ciudad tenga 'San Rafael', pero tiene '%s'", facultad.Ciudad)
		}
	})

	t.Run("Test Domicilio", func(t *testing.T) {
		if facultad.Domicilio != "Av. Gral. J. J. Urquiza 314" {
			t.Errorf("Se esperaba que Domicilio tenga 'Av. Gral. J. J. Urquiza 314', pero tiene '%s'", facultad.Domicilio)
		}
	})

	t.Run("Test Telefono", func(t *testing.T) {
		if facultad.Telefono != "02604421078" {
			t.Errorf("Se esperaba que Telefono tenga '02604421078', pero tiene '%s'", facultad.Telefono)
		}
	})

	t.Run("Test Contacto", func(t *testing.T) {
		if facultad.Contacto != "Ing. Roberto D. Vilches" {
			t.Errorf("Se esperaba que Contacto tenga 'Ing. Roberto D. Vilches', pero tiene '%s'", facultad.Contacto)
		}
	})

	t.Run("Test Email", func(t *testing.T) {
		if facultad.Email != "rvilches@frsr.utn.edu.ar" {
			t.Errorf("Se esperaba que Email tenga 'rvilches@frsr.utn.edu.ar', pero tiene '%s'", facultad.Email)
		}
	})
}
