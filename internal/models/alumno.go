package models

import "time"

type Alumno struct {
	Apellido        string
	Nombre          string
	NroDocumento    string
	TipoDocumento   TipoDocumento
	FechaNacimiento time.Time
	Sexo            string
	NroLegajo       int
	FechaIngreso    time.Time
}
