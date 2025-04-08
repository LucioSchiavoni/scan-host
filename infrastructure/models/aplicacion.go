package models

type Aplicacion struct {
	ID     uint `gorm:"primaryKey"`
	Nombre string
}
