package models

type EquipoAplicacion struct {
	EquipoID     uint
	AplicacionID uint
	Aplicacion   Aplicacion `gorm:"foreignKey:app_id"`
}
