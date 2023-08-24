package main

import (
	"sync"

	"github.com/official-stallion/race/internal/client"
	"github.com/official-stallion/race/internal/config"
)

func main() {
	// load configs
	cfg := config.Load()

	// create consumers
	j := 0
	for i := 0; i < cfg.Consumers; i++ {
		topic := cfg.Topics[j]

		go client.Consumer(cfg.Host, topic)

		j++
		if j == len(cfg.Topics) {
			j = 0
		}
	}

	// create providers
	j = 0
	for i := 0; i < cfg.Providers; i++ {
		topic := cfg.Topics[j]

		go client.Provider(cfg.Host, topic, cfg.Message, cfg.Offset)

		j++
		if j == len(cfg.Topics) {
			j = 0
		}
	}

	var wg sync.WaitGroup

	wg.Add(1)
	wg.Wait()
}
