package models

type EquipoApp struct {
	ID       uint `json:"id" gorm:"primaryKey"`
	IDEquipo uint `json:"id_equipo"`
	IDApp    uint `json:"id_app"`
}
