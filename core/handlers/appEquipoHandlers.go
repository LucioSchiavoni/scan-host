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

func AgregarAplicacionesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	equipoID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "ID de equipo inválido", http.StatusBadRequest)
		return
	}

	var request struct {
		IdApp []uint `json:"id_app"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	for _, appID := range request.IdApp {
		result := usecases.AgregarAplicacion(uint(equipoID), appID)
		if result.Error != "" {
			http.Error(w, result.Error, http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Aplicaciones agregadas correctamente"})
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

func ObtenerAplicacionesHandler(w http.ResponseWriter, r *http.Request) {
	aplicaciones, err := usecases.GetAppService()
	if err != nil {
		http.Error(w, "Error al obtener las aplicaciones", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aplicaciones)
}

func GetAppByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aplicacionID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "ID de aplicación inválido", http.StatusBadRequest)
		return
	}

	aplicacion, err := usecases.GetAppByIdService(uint(aplicacionID))
	if err != nil {
		http.Error(w, "Error al obtener la aplicación", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(aplicacion)
}
