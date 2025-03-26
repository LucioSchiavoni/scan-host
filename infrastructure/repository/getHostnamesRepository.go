package repository

import (
	"github.com/LucioSchiavoni/scan-host/infrastructure/database"
	"github.com/LucioSchiavoni/scan-host/infrastructure/models"
)

func GetHostnamesRepository(hostname string) (*models.Equipo, error) {

	var pc models.Equipo
	err := database.DB.Where("nombre = ?", hostname).First(&pc).Error
	if err != nil {
		return nil, err
	}
	return &pc, nil
}
