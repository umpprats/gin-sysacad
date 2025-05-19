package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

type ServerConfig struct {
	Port string
}

func LoadConfig() (*Config, error) {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	rootDir, err := findProjectRoot()
	if err != nil {
		log.Printf("Advertencia: No se pudo determinar la raíz del proyecto: %v", err)
		viper.AddConfigPath(".")
	} else {
		viper.AddConfigPath(rootDir)
		log.Printf("Buscando archivo de configuración en: %s", rootDir)
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("SYSACAD")

	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.sslmode", "disable")
	viper.SetDefault("database.timezone", "America/Argentina/Buenos_Aires")
	viper.SetDefault("server.port", "8080")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Advertencia: No se pudo leer el archivo de configuración: %v, usando valores predeterminados y variables de entorno", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("no se pudo decodificar la configuración: %v", err)
	}

	fmt.Printf("Configuración cargada correctamente %s", config.Database.DBName)
	return &config, nil
}

func LoadTestConfig() (*Config, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	if !containsTestPrefix(config.Database.DBName) {
		log.Printf("Cambiando de base de datos %s a TEST", config.Database.DBName)
		config.Database.DBName = "TEST_" + config.Database.DBName
	}

	return config, nil
}

// containsTestPrefix verifica si el nombre de la base de datos tiene un prefijo de prueba
func containsTestPrefix(dbName string) bool {
	prefixes := []string{"TEST_", "DEV_", "test_", "dev_"}
	for _, prefix := range prefixes {
		if len(dbName) >= len(prefix) && dbName[:len(prefix)] == prefix {
			return true
		}
	}
	return false
}

func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode, c.TimeZone,
	)
}

func findProjectRoot() (string, error) {
	// Comenzar desde el directorio actual
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Subir en la jerarquía de directorios hasta encontrar go.mod o config.yaml
	for {
		// Comprobar si config.yaml existe en este directorio
		configPath := filepath.Join(dir, "config.yaml")
		if _, err := os.Stat(configPath); err == nil {
			return dir, nil
		}

		// Comprobar si go.mod existe en este directorio
		goModPath := filepath.Join(dir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return dir, nil
		}

		// Ir al directorio padre
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			// Hemos llegado a la raíz del sistema de archivos sin encontrar nada
			break
		}
		dir = parentDir
	}

	return "", fmt.Errorf("no se pudo encontrar la raíz del proyecto")
}
