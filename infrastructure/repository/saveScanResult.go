package repository

import (
	"github.com/LucioSchiavoni/scan-host/infrastructure/database"
	"github.com/LucioSchiavoni/scan-host/infrastructure/models"
)

func SaveScanResult(piso int, ip string, hostname string) {
	result := models.Equipo{
		Nombre: hostname,
		Piso:   piso,
	}

	database.DB.Create(&result)
}
