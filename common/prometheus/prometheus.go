package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

func AddGauge(gaugeName string, gaugeDesc string, applyMetric func(gauge prometheus.Gauge), interval time.Duration) {
	gauge := promauto.NewGauge(prometheus.GaugeOpts{
		Name: gaugeName,
		Help: gaugeDesc,
	})

	go func(gauge prometheus.Gauge) {
		for {
			applyMetric(gauge)
			time.Sleep(interval)
		}
	}(gauge)
}
