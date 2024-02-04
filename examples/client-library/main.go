package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	port := ":8082"

	simpleCounter := promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_simple_count_total",
		Help: "The total number of simple counts",
	})

	// Increment the counter
	simpleCounter.Inc()

	fmt.Printf("listening on port %s\n", port)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(port, nil)
}
