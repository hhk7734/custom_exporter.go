package exporter

import "github.com/prometheus/client_golang/prometheus"

var (
	count = prometheus.NewDesc(
		prometheus.BuildFQName("example", "", "count"),
		"example count",
		nil, nil,
	)
	countVal = 0
)

var _ prometheus.Collector = new(Exporter)

type Exporter struct{}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- count
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.collectExample(ch)
}

func (e *Exporter) collectExample(ch chan<- prometheus.Metric) {
	countVal++
	ch <- prometheus.MustNewConstMetric(count, prometheus.CounterValue, float64(countVal))
}
