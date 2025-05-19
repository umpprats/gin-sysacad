package services_test

import (
	"testing"

	"github.com/umpprats/gin-sysacad/internal/helpers"
	"github.com/umpprats/gin-sysacad/internal/models"
	"github.com/umpprats/gin-sysacad/internal/repositories"
	"github.com/umpprats/gin-sysacad/internal/services"
)

func TestUniversidadService(t *testing.T) {
	// Configurar base de datos de prueba
	db, err := helpers.SetupInTestDB()
	if err != nil {
		t.Fatalf("Error al configurar la base de datos de prueba: %v", err)
	}

	// Crear repositorio y servicio
	repo := repositories.NewUniversidadRepository(db)
	service := services.NewUniversidadService(repo)

	// Datos de prueba
	universidadTest := &models.Universidad{
		Sigla:  "UTN",
		Nombre: "Universidad Tecnológica Nacional",
	}

	// Test CreateUniversidad
	t.Run("Create", func(t *testing.T) {
		created, err := service.Create(universidadTest)
		if err != nil {
			t.Errorf("Error al crear universidad: %v", err)
		}

		tests := []struct {
			nombre   string
			obtenido interface{}
			esperado interface{}
		}{
			{
				nombre:   "ID asignado",
				obtenido: created.ID > 0,
				esperado: true,
			},
			{
				nombre:   "Sigla",
				obtenido: created.Sigla,
				esperado: "UTN",
			},
		}

		for _, test := range tests {
			if test.obtenido != test.esperado {
				t.Errorf("Error en CreateUniversidad - %s: esperado '%v', obtenido '%v'",
					test.nombre, test.esperado, test.obtenido)
			}
		}
	})

	// Test GetUniversidadByID
	t.Run("GetByID", func(t *testing.T) {
		found, err := service.GetByID(universidadTest.ID)
		if err != nil {
			t.Errorf("Error al obtener universidad por ID: %v", err)
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
				t.Errorf("Error en GetUniversidadByID - %s: esperado '%s', obtenido '%s'",
					test.nombre, test.esperado, test.obtenido)
			}
		}
	})

	t.Run("GetAll", func(t *testing.T) {

		otra := &models.Universidad{
			Sigla:  "UBA",
			Nombre: "Universidad de Buenos Aires",
		}
		_, err := service.Create(otra)
		if err != nil {
			t.Fatalf("Error al crear segunda universidad: %v", err)
		}

		universidades, err := service.GetAll()
		if err != nil {
			t.Errorf("Error al obtener todas las universidades: %v", err)
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

		tests := []struct {
			nombre   string
			obtenido bool
			esperado bool
		}{
			{
				nombre:   "Encontró UTN",
				obtenido: encontroUTN,
				esperado: true,
			},
			{
				nombre:   "Encontró UBA",
				obtenido: encontroUBA,
				esperado: true,
			},
		}

		for _, test := range tests {
			if test.obtenido != test.esperado {
				t.Errorf("Error en GetAllUniversidades - %s: esperado '%v', obtenido '%v'",
					test.nombre, test.esperado, test.obtenido)
			}
		}
	})

	t.Run("Update", func(t *testing.T) {

		universidadTest.Nombre = "Universidad Tecnológica Nacional - Actualizada"
		updated, err := service.Update(universidadTest)
		if err != nil {
			t.Errorf("Error al actualizar universidad: %v", err)
		}

		tests := []struct {
			nombre   string
			obtenido string
			esperado string
		}{
			{
				nombre:   "Nombre Actualizado",
				obtenido: updated.Nombre,
				esperado: "Universidad Tecnológica Nacional - Actualizada",
			},
		}

		for _, test := range tests {
			if test.obtenido != test.esperado {
				t.Errorf("Error en UpdateUniversidad - %s: esperado '%s', obtenido '%s'",
					test.nombre, test.esperado, test.obtenido)
			}
		}

		found, err := service.GetByID(universidadTest.ID)
		if err != nil {
			t.Errorf("Error al obtener universidad actualizada: %v", err)
		}

		if found.Nombre != "Universidad Tecnológica Nacional - Actualizada" {
			t.Errorf("El nombre no se actualizó correctamente en la BD. Esperado: '%s', Obtenido: '%s'",
				"Universidad Tecnológica Nacional - Actualizada", found.Nombre)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err := service.Delete(universidadTest.ID)
		if err != nil {
			t.Errorf("Error al eliminar universidad: %v", err)
		}

		_, err = service.GetByID(universidadTest.ID)
		if err == nil {
			t.Error("Se esperaba un error al buscar una universidad eliminada, pero no se produjo")
		}
	})

	t.Run("ValidationRules", func(t *testing.T) {

		invalida := &models.Universidad{
			Sigla:  "",
			Nombre: "Universidad Sin Sigla",
		}

		_, err := service.Create(invalida)
		if err == nil {
			t.Error("Se esperaba un error de validación por sigla vacía, pero no se produjo")
		}
	})
}
