package adapters

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Collector struct {
	DurationHistogram *prometheus.HistogramVec
	ErrorCounter      *prometheus.CounterVec
}

func NewAdapterCollector(namespace string) *Collector {
	return &Collector{
		DurationHistogram: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "contact_repository_request_duration_seconds",
			Help:      "The duration of the requests to contact repository",
		}, []string{"method"}),
		ErrorCounter: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "contact_repository_error_count_total",
			Help:      "The count of the errors to contact repository",
		}, []string{"method"}),
	}
}

func (d *Collector) Describe(desc chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(d, desc)
}

func (d *Collector) Collect(c chan<- prometheus.Metric) {
	d.DurationHistogram.Collect(c)
	d.ErrorCounter.Collect(c)
}
