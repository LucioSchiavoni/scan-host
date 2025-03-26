package repository

import (
	"time"

	"github.com/LucioSchiavoni/scan-host/infrastructure/database"
	"github.com/LucioSchiavoni/scan-host/infrastructure/models"
)

func SaveScanRepository(piso int, hostname string) {
	result := models.Equipo{
		Nombre:      hostname,
		Piso:        piso,
		Estado:      "activo",
		UltimaVista: time.Now(),
	}

	database.DB.Create(&result)
}
