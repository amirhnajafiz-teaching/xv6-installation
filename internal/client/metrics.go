package client

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	Namespace = "stallion_load_test"
	Subsystem = "client"
)

type Metrics struct {
	Consumers          prometheus.Counter
	Providers          prometheus.Counter
	SuccessConnections prometheus.Counter
	FailedConnections  prometheus.Counter
	FailedPublish      prometheus.Counter
}

// newCounter generator.
func newCounter(counterOpts prometheus.CounterOpts) prometheus.Counter {
	ev := prometheus.NewCounter(counterOpts)

	if err := prometheus.Register(ev); err != nil {
		var are prometheus.AlreadyRegisteredError
		if ok := errors.As(err, &are); ok {
			ev, ok = are.ExistingCollector.(prometheus.Counter)
			if !ok {
				panic("different metric type registration")
			}
		} else {
			panic(err)
		}
	}

	return ev
}

func NewMetrics() Metrics {
	return Metrics{
		Consumers: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "number_of_consumers",
			Help:        "total number of consumers",
			ConstLabels: nil,
		}),
		Providers: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "number_of_providers",
			Help:        "total number of providers",
			ConstLabels: nil,
		}),
		SuccessConnections: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "number_of_successful_connections",
			Help:        "total number of connection that were successful",
			ConstLabels: nil,
		}),
		FailedConnections: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "connection_errors_total",
			Help:        "total number of connection errors",
			ConstLabels: nil,
		}),
		FailedPublish: newCounter(prometheus.CounterOpts{
			Namespace:   Namespace,
			Subsystem:   Subsystem,
			Name:        "publish_errors_total",
			Help:        "total number of publish errors",
			ConstLabels: nil,
		}),
	}
}
