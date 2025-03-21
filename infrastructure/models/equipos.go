package models

import (
	"time"
)

type Equipo struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Nombre    string    `json:"nombre"`
	Piso      int       `json:"piso"`
	CreatedAt time.Time `json:"fecha" gorm:"default:CURRENT_TIMESTAMP"`
}
