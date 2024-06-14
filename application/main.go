package main

import (
	"log"
	"net/http"

    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("v0.2.0"))
	})
    http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalf("Failed to start http server. Err-%v", err)
	}

    log.Fatal(http.ListenAndServe(":19100", nil))
}