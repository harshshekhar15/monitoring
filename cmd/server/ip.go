package main

import (
	"fmt"
	"net"
	"net/http"
)

// getIP returns IP address of a client
func getIP(r *http.Request) (string, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("IP address not found")
}
