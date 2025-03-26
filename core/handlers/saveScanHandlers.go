package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LucioSchiavoni/scan-host/core/usecases"
)

func SaveScanHandler(w http.ResponseWriter, r *http.Request) {
	result := usecases.SaveScan()

	w.Header().Set("Content-Type", "application/json")

	if result.Error != "" {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, `{"message": "Error al generar la respuesta JSON"}`, http.StatusInternalServerError)
	}
}
