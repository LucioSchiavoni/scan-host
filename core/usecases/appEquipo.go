package usecases

import (
	"github.com/LucioSchiavoni/scan-host/infrastructure/models"
	"github.com/LucioSchiavoni/scan-host/infrastructure/repository"
)

type EquipoDetalleResult struct {
	Equipo *models.Equipo `json:"equipo"`
	Error  string         `json:"error,omitempty"`
}

type OperacionResult struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func GetEquipoDetalle(equipoID uint) EquipoDetalleResult {
	equipo, err := repository.GetEquipoConAplicaciones(equipoID)
	if err != nil {
		return EquipoDetalleResult{
			Error: err.Error(),
		}
	}
	return EquipoDetalleResult{
		Equipo: equipo,
	}
}

func AgregarAplicacion(equipoID uint, aplicacionID uint) OperacionResult {
	err := repository.AgregarAplicacionAEquipo(equipoID, aplicacionID)
	if err != nil {
		return OperacionResult{
			Message: "Error al agregar la aplicación",
			Error:   err.Error(),
		}
	}
	return OperacionResult{
		Message: "Aplicación agregada correctamente",
	}
}

func RemoverAplicacion(equipoID uint, aplicacionID uint) OperacionResult {
	err := repository.RemoverAplicacionDeEquipo(equipoID, aplicacionID)
	if err != nil {
		return OperacionResult{
			Message: "Error al remover la aplicación",
			Error:   err.Error(),
		}
	}
	return OperacionResult{
		Message: "Aplicación removida correctamente",
	}
}

func GetAppService() ([]*models.Aplicacion, error) {
	return repository.GetApp()
}

func GetAppByIdService(id uint) ([]*models.Aplicacion, error) {
	equipoAplicaciones, err := repository.GetAppsByEquipoID(id)
	if err != nil {
		return nil, err
	}

	var aplicaciones []*models.Aplicacion

	for _, equipoAplicacion := range equipoAplicaciones {
		aplicaciones = append(aplicaciones, &equipoAplicacion.Aplicacion)
	}

	return aplicaciones, nil
}
