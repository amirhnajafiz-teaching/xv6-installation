package config

import (
	"github.com/official-stallion/race/internal/client"
	"github.com/official-stallion/race/internal/telemetry/metrics"
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
