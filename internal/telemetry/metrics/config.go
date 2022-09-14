package metrics

type Config struct {
	Enabled bool   `koanf:"enabled"`
	Port    string `koanf:"port"`
}
