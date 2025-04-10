package models

import "time"

type EquipoAplicacion struct {
	ID               uint       `json:"id" gorm:"primaryKey;table:equipo_aplicacions"`
	EquipoID         uint       `json:"equipo_id"`
	Equipo           Equipo     `json:"equipo" gorm:"foreignKey:EquipoID"`
	AplicacionID     uint       `json:"aplicacion_id"`
	Aplicacion       Aplicacion `json:"aplicacion" gorm:"foreignKey:AplicacionID"`
	FechaInstalacion time.Time  `json:"fecha_instalacion"`
	Estado           string     `json:"estado" gorm:"default:activo"`
}
