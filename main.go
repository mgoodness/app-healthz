package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/mgoodness/app-healthz/healthz"
	"github.com/mgoodness/app-healthz/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

var version = "1.1.0"

func main() {
	log.Println("Starting app...")

	httpAddr := os.Getenv("HTTP_ADDR")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseHost := os.Getenv("DATABASE_HOST")
	databaseName := os.Getenv("DATABASE_NAME")

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Initializing database connection pool...")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		databaseUsername, databasePassword, databaseHost, databaseName)

	hc := &healthz.Config{
		Hostname: hostname,
		Database: healthz.DatabaseConfig{
			DriverName:     "mysql",
			DataSourceName: dataSourceName,
		},
	}

	healthzHandler, err := healthz.Handler(hc)
	if err != nil {
		log.Fatal(err)
	}

	metrics.Register()

	http.Handle("/metrics", prometheus.UninstrumentedHandler())
	http.Handle("/healthz", healthzHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, html, hostname, version)
		metrics.HTTPRequestsTotal.WithLabelValues(
			strconv.Itoa(http.StatusOK), "/", r.Method).Inc()
	})

	log.Printf("HTTP service listening on %s", httpAddr)
	http.ListenAndServe(httpAddr, nil)
}
