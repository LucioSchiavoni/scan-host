package models

import (
	"time"
)

type Equipo struct {
	ID           uint               `json:"id" gorm:"primaryKey;table:equipos"`
	Nombre       string             `json:"nombre"`
	Piso         int                `json:"piso"`
	Estado       string             `json:"estado" gorm:"default:activo"`
	UltimaVista  time.Time          `json:"ultima_vista" gorm:"autoUpdateTime"`
	Aplicaciones []EquipoAplicacion `json:"aplicacion" gorm:"foreignKey:EquipoID"`
}
