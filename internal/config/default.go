package config

// Default configs.
func Default() Config {
	return Config{
		Host:      "",
		Message:   "",
		Debug:     false,
		Offset:    0,
		Consumers: 20,
		Providers: 20,
		Topics:    []string{"stallion"},
	}
}
