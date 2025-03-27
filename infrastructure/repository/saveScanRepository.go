package repository

import (
	"time"

	"github.com/LucioSchiavoni/scan-host/infrastructure/database"
	"github.com/LucioSchiavoni/scan-host/infrastructure/models"
)

func SaveScanRepository(piso int, hostname string) error {
	result := models.Equipo{
		Nombre:      hostname,
		Piso:        piso,
		Estado:      "activo",
		UltimaVista: time.Now(),
	}

	err := database.DB.Create(&result).Error
	if err != nil {
		return err
	}
	return nil

}
