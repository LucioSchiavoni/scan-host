package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LucioSchiavoni/scan-host/config"
	"github.com/LucioSchiavoni/scan-host/core/scans"
	"github.com/gorilla/mux"
)

func ScanAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", config.FrontendDev)
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	results := scans.ScanNetwork(1, 11)
	if results == nil {
		http.Error(w, "Error al escanear la red", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func ScanRange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", config.FrontendDev)
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	vars := mux.Vars(r)
	startSubnet, err1 := strconv.Atoi(vars["startSubnet"])
	endSubnet, err2 := strconv.Atoi(vars["endSubnet"])

	if err1 != nil || err2 != nil || startSubnet > endSubnet {
		http.Error(w, "Parámetros inválidos", http.StatusBadRequest)
		return
	}

	results := scans.ScanNetwork(startSubnet, endSubnet)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
