package consumer

import (
	"sync"

	"github.com/official-stallion/race/internal/client"
	"github.com/official-stallion/race/internal/config"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "consumer",
		Short: "running consumers",
		Run: func(_ *cobra.Command, _ []string) {
			main()
		},
	}
}

func main() {
	// load configs
	cfg := config.Load()

	// load metrics
	metric := client.NewMetrics()

	// creating a wait group
	wg := sync.WaitGroup{}

	for i := 0; i < cfg.Consumers; i++ {
		cli := client.Client{
			Cfg: cfg.Client,
		}

		go func(cli client.Client) {
			wg.Add(1)
			metric.Consumers.Add(1)

			channel := make(chan []byte)

			if err := cli.Connect(); err != nil {
				metric.FailedConnections.Add(1)

				wg.Done()

				return
			}

			metric.SuccessConnections.Add(1)

			cli.Subscribe(cfg.Topic, channel)
		}(cli)
	}

	wg.Wait()
}
