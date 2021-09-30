package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

func ip(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(ClientIP(r)))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	version(w)
	w.WriteHeader(200)
}

func headers(w http.ResponseWriter, req *http.Request) {
	version(w)
	for name, headers := range req.Header {
		for _, h := range headers {
			// fmt.Fprintf(w, "%v: %v\n", name, h)
			fmt.Printf("%v: %v\n", name, h)
			w.Header().Add(name, h)
		}
	}
	ip(w, req)
}

func version(w http.ResponseWriter) {
	w.Header().Add("VERSION", os.Getenv("VERSION"))
}

func main() {
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", headers)
	http.ListenAndServe(":9100", nil)
}
