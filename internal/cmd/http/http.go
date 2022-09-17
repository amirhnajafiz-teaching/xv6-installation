package http

import "github.com/spf13/cobra"

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "starting http command",
		Run: func(_ *cobra.Command, _ []string) {
			main()
		},
	}
}

func main() {

}
