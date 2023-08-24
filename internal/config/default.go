package config

// Default configs.
func Default() Config {
	return Config{
		Consumers: 20,
		Providers: 20,
		Topic:     "stallion",
	}
}
