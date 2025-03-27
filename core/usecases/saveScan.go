package usecases

import (
	"log"
	"time"

	"github.com/LucioSchiavoni/scan-host/core/scans"
	"github.com/LucioSchiavoni/scan-host/infrastructure/repository"
	"gorm.io/gorm"
)

type ScanResult struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func SaveScan() ScanResult {
	log.Println("Iniciando escaneo de red...")
	results := scans.ScanNetwork(1, 11)
	log.Printf("Resultados del escaneo: %+v", results)

	if len(results) == 0 {
		return ScanResult{
			Message: "No se encontraron equipos en la red",
			Error:   "No se detectaron equipos en el rango especificado",
		}
	}

	for _, result := range results {
		log.Printf("Procesando equipo: %s en piso %d", result.Nombre, result.Piso)

		pc, err := repository.GetHostnamesRepository(result.Nombre)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				log.Printf("Equipo no encontrado en DB, intentando guardar: %s", result.Nombre)
				err = repository.SaveScanRepository(result.Piso, result.Nombre)
				if err != nil {
					log.Printf("Error guardando equipo: %v", err)
					return ScanResult{
						Message: "Error guardando nuevo equipo",
						Error:   err.Error(),
					}
				}
				log.Printf("Equipo guardado exitosamente: %s", result.Nombre)
			} else {
				log.Printf("Error consultando DB: %v", err)
				return ScanResult{
					Message: "Error consultando la base de datos",
					Error:   err.Error(),
				}
			}
		} else {
			log.Printf("Equipo encontrado en DB, actualizando: %s", result.Nombre)
			pc.UltimaVista = time.Now()
			if pc.Estado == "inactivo" {
				pc.Estado = "activo"
			}
			err := repository.UpdateHostnamesRepository(pc)
			if err != nil {
				log.Printf("Error actualizando equipo: %v", err)
				return ScanResult{
					Message: "Error actualizando la base de datos",
					Error:   err.Error(),
				}
			}
			log.Printf("Equipo actualizado exitosamente: %s", result.Nombre)
		}
	}

	return ScanResult{Message: "Escaneo completado"}
}
