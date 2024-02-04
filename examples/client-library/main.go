package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	port := ":8082"

	// export a simple counter
	simpleCounter := promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_simple_count_total",
		Help: "The total number of simple counts",
	})

	// Increment the counter
	simpleCounter.Inc()

	// push a simple metric to the pushgateway
	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_completion_timestamp_seconds",
		Help: "The timestamp of the last successful completion of a DB backup.",
	})

	completionTime.SetToCurrentTime()
	if err := push.New("http://localhost:9091", "db_backup").
		Collector(completionTime).
		Grouping("db", "customers").
		Push(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}

	fmt.Printf("listening on port %s\n", port)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(port, nil)
}
