package config

import (
	"github.com/official-stallion/stallion-load-test/internal/client"
	"github.com/official-stallion/stallion-load-test/internal/telemetry/metrics"
)

// Default configs.
func Default() Config {
	return Config{
		Consumers: 20,
		Providers: 20,
		Topic:     "stallion",
		Client: client.Config{
			Address: "",
			Port:    9090,
			Auth: client.Auth{
				Username: "root",
				Password: "Pa$$word",
			},
		},
		Metrics: metrics.Config{
			Enabled: false,
			Port:    "3000",
		},
	}
}
