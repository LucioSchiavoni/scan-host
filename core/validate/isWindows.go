package validate

// func IsWindowsPC(ip string) bool {

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	cmd := exec.Command("nmap", "-p", "445,80", ip)

// 	output, err := cmd.CombinedOutput()

// 	if err != nil {

// 		fmt.Printf("Error ejecutando nmap en %s: %v\n", ip, err)

// 		fmt.Printf("Salida de error: %s\n", string(output))

// 		if ctx.Err() == context.DeadlineExceeded {
// 			fmt.Printf("El comando nmap alcanzó el tiempo límite para la IP %s\n", ip)
// 		}
// 		return false
// 	}

// 	if strings.Contains(string(output), "Windows") {
// 		return true
// 	}

// 	fmt.Printf("No se detectó Windows en la IP %s\n", ip)
// 	return false
// }
