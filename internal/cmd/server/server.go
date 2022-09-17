package server

import "github.com/spf13/cobra"

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "starting server",
		Run: func(_ *cobra.Command, _ []string) {
			main()
		},
	}
}

func main() {

}
