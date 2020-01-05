package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go metricsServer()
	mainServer()
}

func mainServer() {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	port := ":8080"
	fmt.Println("Primary server listening on", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func metricsServer() {
	secondsElapsed := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "seconds_elapsed",
		Help: "Seconds elapsed since this server was started",
	})
	go func() {
		for {
			secondsElapsed.Inc()
			time.Sleep(time.Second)
		}
	}()

	randomNumber := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "random_number",
		Help: "Just a random number [0, 100)",
	}, func() float64 {
		return float64(rand.Intn(100))
	})

	// Create metrics registry
	reg := prometheus.NewRegistry()
	reg.Register(secondsElapsed)
	reg.Register(randomNumber)

	// Create handler that only uses our custom metrics
	handler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		Timeout: time.Second,
	})

	// Serve our metrics
	r := http.NewServeMux()
	r.Handle("/metrics", handler)
	port := ":2112"
	fmt.Println("Metrics server listening on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
