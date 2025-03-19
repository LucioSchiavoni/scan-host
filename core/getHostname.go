package core

import (
	"os/exec"
	"strings"
	"sync"
)

func GetHostname(ip string, mu *sync.Mutex, unknownCount *int, offlineCount *int) string {
	cmd := exec.Command("nslookup", ip)
	output, err := cmd.Output()
	if err != nil {
		mu.Lock()
		*offlineCount++
		mu.Unlock()
		return "Desconocido"
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
	}
	mu.Unlock()

	return domainName
}
