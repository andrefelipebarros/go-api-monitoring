package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total de requisições HTTP",
		},
		[]string{"method", "path"},
	)
)

func main() {
	prometheus.MustRegister(httpRequestsTotal)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		httpRequestsTotal.WithLabelValues(r.Method, "/health").Inc()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	// endpoint de métricas
	http.Handle("/metrics", promhttp.Handler())

	log.Println("API rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
