package usecases

import (
	"github.com/LucioSchiavoni/scan-host/infrastructure/database"
	"github.com/LucioSchiavoni/scan-host/infrastructure/models"
)

type GetScanResult struct {
	Equipos []models.Equipo `json:"equipos"`
	Error   string          `json:"error,omitempty"`
}

func GetScan() GetScanResult {
	var equipos []models.Equipo

	err := database.DB.Find(&equipos).Error
	if err != nil {
		return GetScanResult{
			Error: err.Error(),
		}
	}

	return GetScanResult{
		Equipos: equipos,
	}
}
