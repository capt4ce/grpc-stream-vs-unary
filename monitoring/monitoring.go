package monitoring

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsCounter prometheus.Gauge
)

func StartMonitoring(showZeroRequest bool, metricsPort string) {
	requestsCounter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "test_requests_counter",
		Help: "requests counter",
	})
	prometheus.MustRegister(requestsCounter)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(metricsPort, nil)
	}()
}

func IncrementCounter() {
	requestsCounter.Inc()
}

func DecrementCounter() {
	requestsCounter.Dec()
}
