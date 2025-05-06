package config

import (
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestConnectionDB(t *testing.T) {
	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("Error al cargar la configuración: %v", err)
	}
	dsn := config.Database.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Error, no se puede conectar a la base de datos: %v", err)
	}
	defer sqlDB.Close()
	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("Error ping haciendo ping a la base de datos: %v", err)
	}
	t.Log("Conexión exitosa a la base de datos")

}
