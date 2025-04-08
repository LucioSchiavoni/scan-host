package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LucioSchiavoni/scan-host/infrastructure/repository"
)

type AddAppsRequest struct {
	IDEquipo uint   `json:"id_equipo"`
	AppIDs   []uint `json:"id_app"`
}

func AddAppsToEquipoHandler(w http.ResponseWriter, r *http.Request) {
	var req AddAppsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error al procesar la solicitud", http.StatusBadRequest)
		return
	}

	err := repository.AgregarAplicacionAEquipo(req.IDEquipo, req.AppIDs[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Aplicaciones asociadas correctamente"})
}

// func GetAppsByEquipoHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	idEquipo, err := strconv.Atoi(vars["id_equipo"])
// 	if err != nil {
// 		http.Error(w, "ID de equipo inválido", http.StatusBadRequest)
// 		return
// 	}

// 	apps, err := repository.GetAppsByEquipo(uint(idEquipo))
// 	if err != nil {
// 		http.Error(w, "Error al obtener las aplicaciones del equipo", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(apps)
// }
