package core

import "net"

func GetHostname(ip string) string {
	hostnames, err := net.LookupAddr(ip)
	if err != nil || len(hostnames) == 0 {
		return "Desconocido"
	}
	return hostnames[0]
}
