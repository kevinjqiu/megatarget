package cmd

import "github.com/spf13/cobra"

func newServeCommand() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
