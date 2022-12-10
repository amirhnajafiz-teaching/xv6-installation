package metrics

import (
	"github.com/official-stallion/race/internal/config"
	"github.com/official-stallion/race/internal/telemetry/metrics"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "metrics",
		Short: "starting metrics server",
		Run: func(_ *cobra.Command, _ []string) {
			main()
		},
	}
}

func main() {
	// loading configs
	cfg := config.Load()

	// starting metrics server
	metrics.RunMetricsServer(cfg.Metrics)
}
