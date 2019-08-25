package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tlmiller/janitor/pkg/config"
)

const (
	Config = "config"
	Debug  = "debug"
	DryRun = "dry-run"
)

func BindFlags(c *cobra.Command, v *viper.Viper) {
	v.BindPFlag(config.KeyDryRun, c.Flags().Lookup(DryRun))
}
