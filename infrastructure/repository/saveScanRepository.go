package repository

import (
	"log"
	"time"

	"github.com/LucioSchiavoni/scan-host/infrastructure/database"
	"github.com/LucioSchiavoni/scan-host/infrastructure/models"
)

func SaveScanRepository(piso int, hostname string) error {
	log.Printf("Intentando guardar equipo en DB - Hostname: %s, Piso: %d", hostname, piso)

	result := models.Equipo{
		Nombre:      hostname,
		Piso:        piso,
		Estado:      "activo",
		UltimaVista: time.Now(),
	}

	err := database.DB.Create(&result).Error
	if err != nil {
		log.Printf("Error al guardar en DB: %v", err)
		return err
	}

	log.Printf("Equipo guardado exitosamente en DB - ID: %d", result.ID)
	return nil
}
