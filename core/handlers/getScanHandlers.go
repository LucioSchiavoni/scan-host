package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/LucioSchiavoni/scan-host/core/usecases"
)

func GetScanHandler(w http.ResponseWriter, r *http.Request) {
	result := usecases.GetScan()

	if result.Error != "" {
		http.Error(w, result.Error, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
