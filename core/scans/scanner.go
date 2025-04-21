package scans

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/LucioSchiavoni/scan-host/core"
	"github.com/LucioSchiavoni/scan-host/infrastructure/models"
)

func ScanNetwork(startSubnet, endSubnet int) []models.Equipo {
	var wg sync.WaitGroup
	results := make([]models.Equipo, 0)
	mu := sync.Mutex{}

	baseIP := os.Getenv("BASE_IP")
	if baseIP == "" {
		baseIP = "172.24"
		log.Printf("BASE_IP no configurada, usando valor por defecto: %s", baseIP)
	} else {
		log.Printf("Usando BASE_IP configurada: %s", baseIP)
	}

	log.Printf("Iniciando escaneo desde subnet %d hasta %d", startSubnet, endSubnet)
	hostsEncontrados := 0

	for subnet := startSubnet; subnet <= endSubnet; subnet++ {
		for host := 1; host <= 255; host++ {
			ip := fmt.Sprintf("%s.%d.%d", baseIP, subnet, host)
			wg.Add(1)
			go func(ip string, piso int) {
				defer wg.Done()

				hostname := core.GetHostname(ip)
				if hostname != "Desconocido" {
					log.Printf("Host encontrado - IP: %s, Hostname: %s", ip, hostname)
					hostname = strings.TrimSuffix(hostname, ".mec.local.")
					hostname = strings.TrimRight(hostname, ".")

					mu.Lock()
					results = append(results, models.Equipo{
						Piso:   piso,
						Nombre: hostname,
					})
					hostsEncontrados++
					mu.Unlock()
				}
			}(ip, subnet)
		}
	}

	wg.Wait()
	log.Printf("Escaneo completado. Hosts encontrados: %d", hostsEncontrados)
	return results
}
