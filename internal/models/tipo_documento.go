package models

type TipoDocumento string

const (
	DNI       TipoDocumento = "DNI"
	LC        TipoDocumento = "LC"
	LE        TipoDocumento = "LE"
	CI        TipoDocumento = "CI"
	PASAPORTE TipoDocumento = "PASAPORTE"
)

func (t TipoDocumento) String() string {
	return string(t)
}

func (t TipoDocumento) IsValid() bool {
	valid := false
	switch t {
	case DNI, LC, LE, CI, PASAPORTE:
		valid = true
	}
	return valid
}
