package repository

import (
	"github.com/LucioSchiavoni/scan-host/infrastructure/database"
	"github.com/LucioSchiavoni/scan-host/infrastructure/models"
)

func GetScan() string {
	var equipos []models.Equipo

	err := database.DB.
		Preload("Aplicaciones").
		Preload("Aplicaciones.Aplicacion").
		Find(&equipos).Error

	if err != nil {
		return equipos.New("Error consultando la base de datos", err.Error())
	}

	return equipos
}

func AgregarAplicacionAEquipo(equipoID uint, aplicacionID uint) error {
	equipoApp := models.EquipoAplicacion{
		EquipoID:     equipoID,
		AplicacionID: aplicacionID,
	}

	return database.DB.Create(&equipoApp).Error
}
