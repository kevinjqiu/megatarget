package cmd

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "mt",
	}
	cmd.AddCommand(newServeCommand())
	cmd.AddCommand(newTargetCommand())
	return cmd
}
