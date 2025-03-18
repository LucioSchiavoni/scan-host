package main

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// func isAlive(ip string) bool {
// 	var cmd *exec.Cmd
// 	if runtime.GOOS == "windows" {
// 		cmd = exec.Command("ping", "-n", "1", "-w", "500", ip) // Windows
// 	} else {
// 		cmd = exec.Command("ping", "-c", "1", "-W", "1", ip) // Linux/Mac
// 	}

// 	output, err := cmd.Output()
// 	if err != nil {
// 		return false
// 	}

// 	return strings.Contains(string(output), "TTL=")
// }

func scanIP(ip string, wg *sync.WaitGroup, ch chan<- string, unknownCount *int, offlineCount *int, mu *sync.Mutex) {
	defer wg.Done()

	cmd := exec.Command("nslookup", ip)
	output, err := cmd.Output()
	if err != nil {
		mu.Lock()
		*offlineCount++
		mu.Unlock()
		return
	}

	lines := strings.Split(string(output), "\n")
	var domainName string = "Desconocido"
	for _, line := range lines {
		if strings.Contains(line, "Name:") {
			domainName = strings.TrimSpace(strings.Split(line, ":")[1])
			break
		}
	}

	mu.Lock()
	if domainName == "Desconocido" {
		*unknownCount++
	} else {
		ch <- fmt.Sprintf("| %-15s | %-30s |", ip, domainName)
	}
	mu.Unlock()
}

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

	for {
		fmt.Println("📡 Escaneando la red...")
		fmt.Println("-------------------------------------------------")
		fmt.Println("| IP              | Nombre de Dominio          |")
		fmt.Println("-------------------------------------------------")

		for subnet := startSubNet; subnet <= endSubNet; subnet++ {
			for i := startIP; i <= endIP; i++ {
				ip := fmt.Sprintf("%s%d.%d", baseSubnet, subnet, i)
				wg.Add(1)
				go scanIP(ip, &wg, results, &unknownCount, &offlineCount, &mu)
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

		unknownCount = 0
		offlineCount = 0
		results = make(chan string, (endSubNet-startSubNet+1)*(endIP-startIP+1))

		time.Sleep(60 * time.Second)
	}
}
