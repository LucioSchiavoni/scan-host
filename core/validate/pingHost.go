package validate

import (
	"os/exec"
	"strings"
)

func PingHost(ip string) bool {
	cmd := exec.Command("ping", "-c", "1", "-w", "1", ip)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	return strings.Contains(string(output), "1 packets transmitted, 1 received")
}
