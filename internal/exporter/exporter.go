package exporter

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(&Exporter{})
}

var _ prometheus.Collector = new(Exporter)

type Exporter struct{}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
}
