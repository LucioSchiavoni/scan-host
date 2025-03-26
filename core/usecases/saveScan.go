package usecases

import (
	"fmt"
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

	results := scans.ScanNetwork(1, 11)

	for _, result := range results {

		pc, err := repository.GetHostnamesRepository(result.Nombre)
		if err != nil {

			if err == gorm.ErrRecordNotFound {
				repository.SaveScanRepository(result.Piso, result.Nombre)
			} else {

				fmt.Println("Error consultando la DB:", err)
			}

		} else {
			pc.UltimaVista = time.Now()
			if pc.Estado == "inactivo" {
				pc.Estado = "activo"
			}
			err := repository.UpdateHostnamesRepository(pc)
			if err != nil {
				return ScanResult{
					Message: "Error actualizando la base de datos",
					Error:   err.Error(),
				}
			}

		}
	}

	return ScanResult{Message: "Escaneo completado"}
}
