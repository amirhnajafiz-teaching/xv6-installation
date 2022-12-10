package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RunMetricsServer starts metrics server.
func RunMetricsServer(cfg Config) {
	if cfg.Enabled {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
			panic(err)
		}
	}
}
