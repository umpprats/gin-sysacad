package repositories

import (
	"fmt"

	"github.com/umpprats/gin-sysacad/internal/models"
	"gorm.io/gorm"
)

type UniversidadRepository interface {
	Create(universidad *models.Universidad) error
	FindByID(id uint) (*models.Universidad, error)
	FindAll() ([]*models.Universidad, error)
	Update(universidad *models.Universidad) error
	Delete(id uint) error
}

type universidadRepository struct {
	db *gorm.DB
}

func NewUniversidadRepository(db *gorm.DB) UniversidadRepository {
	return &universidadRepository{db: db}
}

func (r *universidadRepository) Create(universidad *models.Universidad) error {
	return r.db.Create(universidad).Error
}

func (r *universidadRepository) FindByID(id uint) (*models.Universidad, error) {
	var universidad models.Universidad
	result := r.db.First(&universidad, id)
	if result.Error != nil {
		return nil, fmt.Errorf("universidad no encontrada: %v", result.Error)
	}
	return &universidad, nil
}

func (r *universidadRepository) FindAll() ([]*models.Universidad, error) {
	var universidades []*models.Universidad
	result := r.db.Find(&universidades)
	if result.Error != nil {
		return nil, result.Error
	}
	return universidades, nil
}

func (r *universidadRepository) Update(universidad *models.Universidad) error {
	return r.db.Save(universidad).Error
}

func (r *universidadRepository) Delete(id uint) error {
	return r.db.Delete(&models.Universidad{}, id).Error
}
