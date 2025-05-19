package helpers

import (
	"fmt"
	"log"

	"github.com/umpprats/gin-sysacad/config"
	"github.com/umpprats/gin-sysacad/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupInTestDB() (*gorm.DB, error) {
	config, err := config.LoadTestConfig()

	if err != nil {
		return nil, fmt.Errorf("error al cargar configuración: %v", err)
	}

	dsn := config.Database.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	migrateModels(db)
	CleanupDB(db)

	return db, nil
}

func migrateModels(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Universidad{},
		&models.Facultad{},
		&models.Alumno{},
	)
	if err != nil {
		return fmt.Errorf("error migrando modelos: %v", err)
	}
	return nil
}

func CleanupDB(db *gorm.DB) {

	tables := []string{"universidades", "facultades", "alumnos"}

	for _, table := range tables {
		if err := CleanupTable(db, table); err != nil {
			log.Printf("Error al limpiar tabla %s: %v", table, err)
		}
	}
}

func CleanupTable(db *gorm.DB, tableName string) error {

	if db == nil {
		return fmt.Errorf("la conexión a la base de datos no está disponible")
	}

	var count int64
	if err := db.Table(tableName).Count(&count).Error; err != nil {
		return fmt.Errorf("error al contar filas en la tabla %s: %v", tableName, err)
	}
	if count == 0 {
		return fmt.Errorf("la tabla %s no existe o está vacía", tableName)
	}

	result := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tableName))
	return result.Error
}
