package config

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/official-stallion/race/internal/client"
	"github.com/official-stallion/race/internal/telemetry/metrics"
)

// Config stores data configs for load test.
type Config struct {
	Consumers int            `koanf:"consumers"`
	Providers int            `koanf:"providers"`
	Topic     string         `koanf:"topic"`
	Client    client.Config  `koanf:"client"`
	Metrics   metrics.Config `koanf:"metrics"`
}

// Load configs.
func Load() Config {
	var instance Config

	k := koanf.New(".")

	// load default
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading deafult: %v\n", err)
	}

	// load configs file
	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config.yaml file: %v\n", err)
	}

	// unmarshalling
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %v\n", err)
	}

	return instance
}
