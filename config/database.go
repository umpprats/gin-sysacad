package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() (*gorm.DB, error) {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error al cargar la configuración: %v", err)
		return nil, err
	}
	dsn := config.Database.GetDSN()

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos")
	}
	return DB, err
}
