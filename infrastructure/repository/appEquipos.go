package repository

import (
	"errors"
	"log"
	"time"

	"github.com/LucioSchiavoni/scan-host/infrastructure/database"
	"github.com/LucioSchiavoni/scan-host/infrastructure/models"
)

func GetEquipoConAplicaciones(equipoID uint) (*models.Equipo, error) {
	var equipo models.Equipo
	err := database.DB.Preload("Aplicaciones.Aplicacion").First(&equipo, equipoID).Error
	if err != nil {
		return nil, err
	}
	return &equipo, nil
}

func AgregarAplicacionAEquipo(equipoID uint, aplicacionID uint) error {

	var aplicacion models.Aplicacion
	if err := database.DB.First(&aplicacion, aplicacionID).Error; err != nil {
		log.Printf("Aplicación no encontrada: ID %d", aplicacionID)
		return errors.New("la aplicación no existe")
	}

	var equipo models.Equipo
	if err := database.DB.First(&equipo, equipoID).Error; err != nil {
		return errors.New("el equipo no existe")
	}

	var equipoApp models.EquipoAplicacion
	if err := database.DB.Where("equipo_id = ? AND aplicacion_id = ?", equipoID, aplicacionID).First(&equipoApp).Error; err == nil {
		return database.DB.Model(&equipoApp).Update("estado", "activo").Error
	}

	equipoApp = models.EquipoAplicacion{
		EquipoID:         equipoID,
		AplicacionID:     aplicacionID,
		FechaInstalacion: time.Now(),
		Estado:           "activo",
	}
	return database.DB.Create(&equipoApp).Error
}

func RemoverAplicacionDeEquipo(equipoID uint, aplicacionID uint) error {
	return database.DB.Model(&models.EquipoAplicacion{}).
		Where("equipo_id = ? AND aplicacion_id = ?", equipoID, aplicacionID).
		Update("estado", "inactivo").Error
}

func ActualizarEquipo(equipo *models.Equipo) error {
	return database.DB.Save(equipo).Error
}

func GetApp() ([]*models.Aplicacion, error) {
	var aplicaciones []models.Aplicacion
	if err := database.DB.Find(&aplicaciones).Error; err != nil {
		return nil, err
	}
	var aplicacionesPtr []*models.Aplicacion
	for i := range aplicaciones {
		aplicacionesPtr = append(aplicacionesPtr, &aplicaciones[i])
	}
	return aplicacionesPtr, nil
}

func GetAppsByEquipoID(equipoID uint) ([]models.EquipoAplicacion, error) {
	var equipoAplicaciones []models.EquipoAplicacion

	err := database.DB.
		Preload("Aplicacion").
		Where("equipo_id = ?", equipoID).
		Find(&equipoAplicaciones).Error

	if err != nil {
		return nil, err
	}

	return equipoAplicaciones, nil
}
