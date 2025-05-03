package models

type Autoridad struct {
	Nombre   string
	Cargo    *Cargo
	Telefono string
	Email    string
}

type Cargo struct {
	Nombre string
	Puntos int
}
