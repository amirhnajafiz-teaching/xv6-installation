package server

import (
	"github.com/official-stallion/stallion"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "starting stallion server",
		Run: func(_ *cobra.Command, _ []string) {
			main()
		},
	}
}

func main() {
	if err := stallion.NewServer(":9254", "root", "Pa$$word"); err != nil {
		panic(err)
	}
}
