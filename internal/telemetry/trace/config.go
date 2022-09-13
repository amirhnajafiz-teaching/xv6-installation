package trace

type Config struct {
	Host  string  `koanf:"host"`
	Port  string  `koanf:"port"`
	Ratio float64 `koanf:"ratio"`
}
