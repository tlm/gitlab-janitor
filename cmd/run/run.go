package run

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/tlmiller/gitlab-janitor/cmd/flags"
	"github.com/tlmiller/gitlab-janitor/pkg/config"
	"github.com/tlmiller/gitlab-janitor/pkg/gitlab/backup"
)

func NewCmdRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "run",
		Short:         "run janitor job once",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE:          run,
	}
	return cmd
}

func run(cmd *cobra.Command, a []string) error {
	builder := config.NewBuilder()
	flags.BindFlags(cmd, builder.Viper)
	confFile, err := cmd.Flags().GetString(flags.Config)
	if err != nil {
		return fmt.Errorf("missing flag config: %v", err)
	}

	var (
		conf *config.Config
	)
	if confFile != "" {
		conf, err = builder.BuildWithConfFile(confFile)
	} else {
		conf, err = builder.Build()
	}

	if err != nil {
		return fmt.Errorf("building config: %v", err)
	}

	bucket, err := config.ToBucket(conf.Bucket)
	if err != nil {
		return fmt.Errorf("getting backup bucket: %v", err)
	}

	decider, err := config.ToDecider(conf.Decider)
	if err != nil {
		return fmt.Errorf("getting backup decider: %v", err)
	}

	pruneList, err := backup.CreatePruneList(bucket, decider)
	if err != nil {
		return fmt.Errorf("generating backup prune list: %v", err)
	}

	if !conf.DryRun {
		log.Printf("deleting %d backups", len(pruneList))
		if err := backup.DeletePruneList(bucket, pruneList); err != nil {
			return fmt.Errorf("deleting backups: %v", err)
		}
	}

	defer bucket.Close()
	return nil
}
