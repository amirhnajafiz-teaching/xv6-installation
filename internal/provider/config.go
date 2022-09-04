package provider

type Config struct {
	Topics []string `koanf:"topics"`
	Number int      `koanf:"number"`
}
