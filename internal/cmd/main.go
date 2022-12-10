package cmd

import (
	"github.com/official-stallion/race/internal/cmd/consumer"
	"github.com/official-stallion/race/internal/cmd/metrics"
	"github.com/official-stallion/race/internal/cmd/publisher"
	"github.com/official-stallion/race/internal/cmd/server"
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(
		consumer.GetCommand(),
		publisher.GetCommand(),
		server.GetCommand(),
		metrics.GetCommand(),
	)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
