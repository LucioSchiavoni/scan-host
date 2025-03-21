package scans

import (
	"fmt"
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
	}

	for subnet := startSubnet; subnet <= endSubnet; subnet++ {
		for host := 1; host <= 255; host++ {
			ip := fmt.Sprintf("%s.%d.%d", baseIP, subnet, host)
			wg.Add(1)
			go func(ip string, piso int) {
				defer wg.Done()
				hostname := core.GetHostname(ip)

				hostname = strings.TrimSuffix(hostname, ".mec.local.")
				hostname = strings.TrimRight(hostname, ".")

				if hostname != "Desconocido" {
					mu.Lock()
					results = append(results, models.Equipo{
						Piso:   piso,
						Nombre: hostname,
					})
					mu.Unlock()
				}
			}(ip, subnet)

		}
	}

	wg.Wait()
	return results
}
