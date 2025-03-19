package main

import (
	"fmt"
	"sync"

	"github.com/LucioSchiavoni/scan-host/core"
)

func main() {
	startSubNet := 9
	endSubNet := 9
	startIP := 1
	endIP := 255
	baseSubnet := "172.24."

	var wg sync.WaitGroup
	results := make(chan string, (endSubNet-startSubNet+1)*(endIP-startIP+1))
	var unknownCount, offlineCount int
	var mu sync.Mutex

	fmt.Println("📡 Escaneando la red...")
	fmt.Println("-------------------------------------------------")
	fmt.Println("| IP              | Nombre de Dominio          |")
	fmt.Println("-------------------------------------------------")

	for subnet := startSubNet; subnet <= endSubNet; subnet++ {
		for i := startIP; i <= endIP; i++ {
			ip := fmt.Sprintf("%s%d.%d", baseSubnet, subnet, i)

			// if !validate.PingHost(ip) {
			// 	continue
			// }

			wg.Add(1)
			go func(ip string) {
				defer wg.Done()

				hostname := core.GetHostname(ip, &mu, &unknownCount, &offlineCount)

				mu.Lock()
				if hostname != "Desconocido" {
					results <- fmt.Sprintf("| %-15s | %-30s |", ip, hostname)
				}
				mu.Unlock()
			}(ip)
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("-------------------------------------------------")
	fmt.Printf("🔍 IPs desconocidas: %d\n", unknownCount)
	fmt.Printf("❌ IPs apagadas: %d\n", offlineCount)
}
