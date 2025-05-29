package exporter

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "custom"
)

var (
	durationDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "scrape", "duration_seconds"),
		"Returns how long the scrape took to complete in seconds.",
		nil,
		nil,
	)
)

var _ prometheus.Collector = (*Exporter)(nil)

type Exporter struct{}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- durationDesc
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		ch <- prometheus.MustNewConstMetric(
			durationDesc,
			prometheus.GaugeValue,
			duration,
		)
	}()

	// Collect metrics
}
