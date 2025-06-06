package models

import (
	"testing"
)

func TestUniversidad(t *testing.T) {

	universidad := Universidad{
		Sigla:  "UTN",
		Nombre: "Universidad Tecnológica Nacional",
	}

	t.Run("Test Sigla", func(t *testing.T) {
		if universidad.Sigla != "UTN" {
			t.Errorf("Se esperaba que la propiedad Sigla tenga 'UTN', pero tiene '%s'", universidad.Sigla)
		}
	})

	t.Run("Test Nombre", func(t *testing.T) {
		if universidad.Nombre != "Universidad Tecnológica Nacional" {
			t.Errorf("Se esperaba que la propiedad Nombre tenga 'Universidad Tecnológica Nacional', pero tiene '%s'", universidad.Nombre)
		}
	})

	t.Run("Test TableName", func(t *testing.T) {
		if universidad.TableName() != "universidades" {
			t.Errorf("Se esperaba que el nombre de la tabla sea 'universidades', pero es '%s'", universidad.TableName())
		}
	})

}
