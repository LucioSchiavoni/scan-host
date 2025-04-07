package repository

import (
	"errors"

	"github.com/LucioSchiavoni/scan-host/infrastructure/database"
	"github.com/LucioSchiavoni/scan-host/infrastructure/models"
)

func AddAppsToEquipo(idEquipo uint, appIDs []uint) error {
	var equipo models.Equipo
	if err := database.DB.First(&equipo, idEquipo).Error; err != nil {
		return errors.New("Equipo no encontrado")
	}

	for _, appID := range appIDs {
		var app models.Aplicacion
		if err := database.DB.First(&app, appID).Error; err != nil {
			return errors.New("una de las aplicaciones no existe")
		}

		equipoApp := models.EquipoApp{
			IDEquipo: idEquipo,
			IDApp:    appID,
		}
		if err := database.DB.Create(&equipoApp).Error; err != nil {
			return errors.New("error al asociar la aplicación al equipo")
		}
	}

	return nil

}

func GetAppsByEquipo(idEquipo uint) ([]models.Aplicacion, error) {
	var equipo models.Equipo
	err := database.DB.Preload("aplicacion").First(&equipo, idEquipo).Error
	if err != nil {
		return nil, err
	}
	return equipo.Aplicaciones, nil
}
