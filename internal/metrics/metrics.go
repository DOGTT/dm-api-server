package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	MetricsTest prometheus.Counter
)

func Init() {
	MetricsTest = promauto.NewCounter(prometheus.CounterOpts{
		Name: "z_test",
		Help: "The total number of processed events",
	})

}
