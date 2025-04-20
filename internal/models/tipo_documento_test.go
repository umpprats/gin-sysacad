package models

import "testing"

func TestTipoDocumento(t *testing.T) {
	tests := []struct {
		tipoDoc  TipoDocumento
		esperado string
	}{
		{DNI, "DNI"},
		{LC, "LC"},
		{LE, "LE"},
		{CI, "CI"},
		{PASAPORTE, "PASAPORTE"},
	}

	for _, test := range tests {
		t.Run(string(test.tipoDoc), func(t *testing.T) {
			resultado := test.tipoDoc.String()
			if resultado != test.esperado {
				t.Errorf("Para %v se esperaba %s pero se obtuvo %s",
					test.tipoDoc, test.esperado, resultado)
			}
		})
	}
}

func TestTipoDocumentoIsValid(t *testing.T) {

	tiposValidos := []TipoDocumento{DNI, LC, LE, CI, PASAPORTE}
	for _, tipo := range tiposValidos {
		t.Run(string(tipo), func(t *testing.T) {
			if !tipo.IsValid() {
				t.Errorf("TipoDocumento %s debería ser válido", tipo)
			}
		})
	}

	tiposInvalidos := []TipoDocumento{"", "OTRO", "CEDULA"}
	for _, tipo := range tiposInvalidos {
		t.Run(string(tipo), func(t *testing.T) {
			if tipo.IsValid() {
				t.Errorf("TipoDocumento %s debería ser válido", tipo)
			}
		})
	}
}
