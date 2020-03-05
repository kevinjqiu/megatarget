package cmd

import "github.com/spf13/cobra"

func newTargetCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:     "target",
		Aliases: []string{"t"},
	}

	cmd.AddCommand(&cobra.Command{
		Use: "ls",
		Run: func(cmd *cobra.Command, args []string) {

		},
	})

	cmd.AddCommand(&cobra.Command{
		Use: "new",
		Run: func(cmd *cobra.Command, args []string) {

		},
	})
	return &cmd
}
