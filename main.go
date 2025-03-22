package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LucioSchiavoni/scan-host/config"
	"github.com/LucioSchiavoni/scan-host/core/handlers"

	"github.com/gorilla/mux"
)

func main() {

	config.LoadConfig()
	// database.ConnectDB()

	r := mux.NewRouter()
	r.HandleFunc("/scan", handlers.ScanAll).Methods("GET")
	r.HandleFunc("/scan/{startSubnet}/{endSubnet}", handlers.ScanRange).Methods("GET")

	serverAddress := fmt.Sprintf(":%s", config.ServerPort)
	log.Printf("🚀 Servidor corriendo en http://localhost%s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, r))

}
