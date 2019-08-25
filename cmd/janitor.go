package cmd

import (
	"github.com/spf13/cobra"

	"github.com/tlmiller/janitor/cmd/flags"
	"github.com/tlmiller/janitor/cmd/run"
)

func New() (command *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "janitor",
		Short: "gitlab backup janitor",
	}
	cmd.PersistentFlags().StringP(flags.Config, "c", "", "configuration file")
	cmd.PersistentFlags().Bool(flags.DryRun, false, "dry run mode, no data is deleted")
	cmd.AddCommand(run.NewCmdRun())
	return cmd
}
