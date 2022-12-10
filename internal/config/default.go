package config

import (
	"github.com/official-stallion/stallion-load-test/internal/client"
)

// Default configs.
func Default() Config {
	return Config{
		Client: client.Config{
			Address: "",
			Port:    9090,
			Auth: client.Auth{
				Username: "root",
				Password: "Pa$$word",
			},
		},
	}
}
