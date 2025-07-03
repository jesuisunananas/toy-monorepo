package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go func() {
		log.Println("pprof listening on :6060")
		http.ListenAndServe("localhost:6060", nil)
	}()

	http.Handle("/metrics", promhttp.Handler())
	go startClient() // Periodically push/load metrics
	go func() {
		time.Sleep(1 * time.Minute)
		runtime.GC()
		log.Println("forced GC")
	}()

	log.Println("Serving metrics on :2112")
	http.ListenAndServe(":2112", nil)
}
