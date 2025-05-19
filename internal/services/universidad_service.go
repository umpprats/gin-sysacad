package services

import (
	"fmt"

	"github.com/umpprats/gin-sysacad/internal/models"
	"github.com/umpprats/gin-sysacad/internal/repositories"
)

type UniversidadService interface {
	Create(universidad *models.Universidad) (*models.Universidad, error)
	GetByID(id uint) (*models.Universidad, error)
	GetAll() ([]*models.Universidad, error)
	Update(universidad *models.Universidad) (*models.Universidad, error)
	Delete(id uint) error
}

type universidadService struct {
	repo repositories.UniversidadRepository
}

func NewUniversidadService(repo repositories.UniversidadRepository) UniversidadService {
	return &universidadService{repo: repo}
}

func (s *universidadService) Create(universidad *models.Universidad) (*models.Universidad, error) {
	if err := s.validate(universidad); err != nil {
		return nil, err
	}

	err := s.repo.Create(universidad)
	if err != nil {
		return nil, fmt.Errorf("error al crear universidad: %v", err)
	}

	return universidad, nil
}

func (s *universidadService) GetByID(id uint) (*models.Universidad, error) {
	return s.repo.FindByID(id)
}

func (s *universidadService) GetAll() ([]*models.Universidad, error) {
	return s.repo.FindAll()
}

func (s *universidadService) Update(universidad *models.Universidad) (*models.Universidad, error) {
	if err := s.validate(universidad); err != nil {
		return nil, err
	}

	_, err := s.repo.FindByID(universidad.ID)
	if err != nil {
		return nil, fmt.Errorf("universidad no encontrada: %v", err)
	}

	err = s.repo.Update(universidad)
	if err != nil {
		return nil, fmt.Errorf("error al actualizar universidad: %v", err)
	}

	return universidad, nil
}

func (s *universidadService) Delete(id uint) error {

	_, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("universidad no encontrada: %v", err)
	}

	return s.repo.Delete(id)
}

func (s *universidadService) validate(universidad *models.Universidad) error {
	if universidad.Sigla == "" {
		return fmt.Errorf("la sigla de la universidad no puede estar vacía")
	}
	if universidad.Nombre == "" {
		return fmt.Errorf("el nombre de la universidad no puede estar vacío")
	}
	return nil
}
