package models

type Aplicacion struct {
	ID      uint   `json:"id" gorm:"primaryKey;table:aplicacions"`
	Nombre  string `json:"nombre"`
	Version string `json:"version"`
}

func (Aplicacion) TableName() string {
	return "aplicacions"
}
