package metrics

// Config stores information about metrics server.
type Config struct {
	Enabled bool   `koanf:"enabled"`
	Port    string `koanf:"port"`
}
