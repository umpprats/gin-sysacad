package repositories_test

import (
	"testing"

	"github.com/umpprats/gin-sysacad/internal/helpers"
	"github.com/umpprats/gin-sysacad/internal/models"
	"github.com/umpprats/gin-sysacad/internal/repositories"
)

func TestUniversidadRepository(t *testing.T) {

	db, err := helpers.SetupInTestDB()
	if err != nil {
		t.Fatalf("Error al configurar la base de datos de prueba: %v", err)
	}

	repo := repositories.NewUniversidadRepository(db)

	universidad := &models.Universidad{
		Sigla:  "UTN",
		Nombre: "Universidad Tecnológica Nacional",
	}

	t.Run("Create", func(t *testing.T) {
		err := repo.Create(universidad)
		if err != nil {
			t.Errorf("Error al crear universidad: %v", err)
		}

		if universidad.ID == 0 {
			t.Error("No se asignó un ID después de crear la universidad")
		}
	})

	t.Run("FindByID", func(t *testing.T) {
		found, err := repo.FindByID(universidad.ID)
		if err != nil {
			t.Errorf("Error al buscar universidad por ID: %v", err)
		}

		tests := []struct {
			nombre   string
			obtenido string
			esperado string
		}{
			{
				nombre:   "Sigla",
				obtenido: found.Sigla,
				esperado: "UTN",
			},
			{
				nombre:   "Nombre",
				obtenido: found.Nombre,
				esperado: "Universidad Tecnológica Nacional",
			},
		}

		for _, test := range tests {
			if test.obtenido != test.esperado {
				t.Errorf("Error en FindByID - %s: esperado '%s', obtenido '%s'",
					test.nombre, test.esperado, test.obtenido)
			}
		}
	})

	t.Run("FindAll", func(t *testing.T) {

		otra := &models.Universidad{
			Sigla:  "UBA",
			Nombre: "Universidad de Buenos Aires",
		}
		err := repo.Create(otra)
		if err != nil {
			t.Fatalf("Error al crear segunda universidad: %v", err)
		}

		universidades, err := repo.FindAll()
		if err != nil {
			t.Errorf("Error al buscar todas las universidades: %v", err)
		}

		if len(universidades) < 2 {
			t.Errorf("Se esperaban al menos 2 universidades, pero se obtuvieron %d", len(universidades))
		}

		var encontroUTN, encontroUBA bool
		for _, u := range universidades {
			if u.Sigla == "UTN" {
				encontroUTN = true
			}
			if u.Sigla == "UBA" {
				encontroUBA = true
			}
		}

		if !encontroUTN {
			t.Error("No se encontró la universidad UTN en los resultados")
		}

		if !encontroUBA {
			t.Error("No se encontró la universidad UBA en los resultados")
		}
	})

	t.Run("Update", func(t *testing.T) {

		universidad.Nombre = "Universidad Tecnológica Nacional (Actualizada)"
		err := repo.Update(universidad)
		if err != nil {
			t.Errorf("Error al actualizar universidad: %v", err)
		}

		updated, err := repo.FindByID(universidad.ID)
		if err != nil {
			t.Errorf("Error al buscar universidad actualizada: %v", err)
		}

		tests := []struct {
			nombre   string
			obtenido string
			esperado string
		}{
			{
				nombre:   "Nombre Actualizado",
				obtenido: updated.Nombre,
				esperado: "Universidad Tecnológica Nacional (Actualizada)",
			},
		}

		for _, test := range tests {
			if test.obtenido != test.esperado {
				t.Errorf("Error en Update - %s: esperado '%s', obtenido '%s'",
					test.nombre, test.esperado, test.obtenido)
			}
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err := repo.Delete(universidad.ID)
		if err != nil {
			t.Errorf("Error al eliminar universidad: %v", err)
		}

		_, err = repo.FindByID(universidad.ID)
		if err == nil {
			t.Error("Se esperaba un error al buscar una universidad eliminada, pero no se produjo")
		}
	})
}
