package models

type Aplicacion struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Nombre  string `json:"nombre"`
	Version string `json:"version"`
}
