package config

// Default configs.
func Default() Config {
	return Config{
		Host:      "",
		Consumers: 20,
		Providers: 20,
		Topics:    []string{"stallion"},
	}
}
