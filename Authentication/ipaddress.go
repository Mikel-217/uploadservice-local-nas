package authentication

import (
	"net/http"
	"strings"
)

func GetIP(r *http.Request) string {

	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}

	if xRealIP := r.Header.Get("X-Real-Ip"); xRealIP != "" {
		return strings.TrimSpace(xRealIP)
	}

	ipPort := r.RemoteAddr
	ip := ipPort
	if strings.Contains(ipPort, ":") {
		ip = strings.Split(ipPort, ":")[0]
	}

	return ip
}
