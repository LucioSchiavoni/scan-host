package models

import (
	"time"
)

type Equipo struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Nombre      string    `json:"nombre"`
	Piso        int       `json:"piso"`
	Estado      string    `gorm:"default:activo"`
	UltimaVista time.Time `gorm:"autoUpdateTime"`
}
