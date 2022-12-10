package publisher

import (
	"sync"
	"time"

	"github.com/official-stallion/race/internal/client"
	"github.com/official-stallion/race/internal/config"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "publisher",
		Short: "running publisher command",
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
			metric.Providers.Add(1)

			if err := cli.Connect(); err != nil {
				metric.FailedConnections.Add(1)

				wg.Done()

				return
			}

			metric.SuccessConnections.Add(1)

			for {
				if err := cli.Publish(cfg.Topic, []byte("message")); err != nil {
					metric.FailedPublish.Add(1)
				}

				time.Sleep(3 * time.Second)
			}
		}(cli)
	}

	wg.Wait()
}
