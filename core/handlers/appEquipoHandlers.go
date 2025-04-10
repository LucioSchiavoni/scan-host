package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LucioSchiavoni/scan-host/core/usecases"
	"github.com/gorilla/mux"
)

func GetEquipoDetalleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	equipoID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "ID de equipo inválido", http.StatusBadRequest)
		return
	}

	result := usecases.GetEquipoDetalle(uint(equipoID))
	if result.Error != "" {
		http.Error(w, result.Error, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func AgregarAplicacionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	equipoID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "ID de equipo inválido", http.StatusBadRequest)
		return
	}

	var request struct {
		AplicacionID uint `json:"aplicacion_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	result := usecases.AgregarAplicacion(uint(equipoID), request.AplicacionID)
	if result.Error != "" {
		http.Error(w, result.Error, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func RemoverAplicacionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	equipoID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "ID de equipo inválido", http.StatusBadRequest)
		return
	}

	aplicacionID, err := strconv.ParseUint(vars["aplicacionId"], 10, 32)
	if err != nil {
		http.Error(w, "ID de aplicación inválido", http.StatusBadRequest)
		return
	}

	result := usecases.RemoverAplicacion(uint(equipoID), uint(aplicacionID))
	if result.Error != "" {
		http.Error(w, result.Error, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
