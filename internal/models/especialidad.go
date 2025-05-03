package models

type Especialidad struct {
	Nombre           string
	Letra            string
	Observacion      string
	TipoEspecialidad *TipoEspecialidad
}

type TipoEspecialidad struct {
	Nombre string
	Nivel  string
}
