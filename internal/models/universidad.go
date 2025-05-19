package models

type Universidad struct {
	ID     uint   `gorm:"primaryKey;column:id; type:bigint;not null"`
	Nombre string `gorm:"column:nombre; type:varchar(100);not null"`
	Sigla  string `gorm:"column:sigla; unique; type:varchar(5);not null"`
}

func (Universidad) TableName() string {
	return "universidades"
}
