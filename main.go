package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LucioSchiavoni/scan-host/config"
	"github.com/LucioSchiavoni/scan-host/core/handlers"
	"github.com/LucioSchiavoni/scan-host/core/middleware"
	"github.com/LucioSchiavoni/scan-host/infrastructure/database"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	database.ConnectDB()

	r := mux.NewRouter()

	// Aplicar el middleware CORS a todas las rutas
	r.Use(middleware.CORSMiddleware)

	// Configurar las rutas
	r.HandleFunc("/scans", handlers.GetScanHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/saveScan", handlers.SaveScanHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/scan/{startSubnet}/{endSubnet}", handlers.ScanRange).Methods("GET", "OPTIONS")
	r.HandleFunc("/equipos/{id}", handlers.GetEquipoDetalleHandler).Methods("GET", "OPTIONS")
	r.HandleFunc("/equipos/{id}/aplicaciones", handlers.AgregarAplicacionesHandler).Methods("PUT", "OPTIONS")
	r.HandleFunc("/equipos/{id}/aplicaciones/{aplicacionId}", handlers.RemoverAplicacionHandler).Methods("DELETE", "OPTIONS")

	serverAddress := fmt.Sprintf(":%s", config.ServerPort)
	log.Printf("🚀 Servidor corriendo en http://localhost%s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, r))
}
