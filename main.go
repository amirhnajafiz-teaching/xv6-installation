package main

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/official-stallion/race/internal/client"
	"github.com/official-stallion/race/internal/config"
	"github.com/official-stallion/race/internal/metrics"
)

func handler(metrics *metrics.Handler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, _ *http.Request) {
		js, err := json.Marshal(metrics.Metrics)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)

			return
		}

		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write(js)
	}
}

func main() {
	// load configs
	cfg := config.Load()

	// create new client
	cli := client.Client{
		Metrics: metrics.New(cfg.Providers, cfg.Consumers, cfg.Topics),
	}

	// create consumers
	j := 0
	for i := 0; i < cfg.Consumers; i++ {
		topic := cfg.Topics[j]

		go cli.Consumer(cfg.Host, topic)

		j++
		if j == len(cfg.Topics) {
			j = 0
		}
	}

	// create providers
	j = 0
	for i := 0; i < cfg.Providers; i++ {
		topic := cfg.Topics[j]

		go cli.Provider(cfg.Host, topic, cfg.Message, cfg.Offset)

		j++
		if j == len(cfg.Topics) {
			j = 0
		}
	}

	var wg sync.WaitGroup

	wg.Add(1)
	wg.Wait()
}
