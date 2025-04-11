package models

type Aplicacion struct {
	ID      uint   `json:"id" gorm:"primaryKey;table:aplicacion"`
	Nombre  string `json:"nombre"`
	Version string `json:"version"`
}

func (Aplicacion) TableName() string {
	return "aplicacion"
}
