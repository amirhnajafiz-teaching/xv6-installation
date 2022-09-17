package publisher

import "github.com/spf13/cobra"

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

}
