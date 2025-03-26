package repository

import (
	"github.com/LucioSchiavoni/scan-host/infrastructure/database"
	"github.com/LucioSchiavoni/scan-host/infrastructure/models"
)

func UpdateHostnamesRepository(pc *models.Equipo) error {

	return database.DB.Save(pc).Error
}
