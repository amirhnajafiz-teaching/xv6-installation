package provider

type Config struct {
	Number  int      `koanf:"number"`
	Timeout int      `koanf:"timeout"`
	Topics  []string `koanf:"topics"`
}
