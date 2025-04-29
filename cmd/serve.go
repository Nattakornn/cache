package cmd

import (
	"github.com/Nattakornn/cache/config"
	"github.com/Nattakornn/cache/modules/servers"
	"github.com/Nattakornn/cache/pkg/logger"
	"github.com/spf13/cobra"
)

var ComposerCoreCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Cache Service",
	RunE: func(cmd *cobra.Command, args []string) error {

		cfg := config.LoadConfig()

		logger.InitZapLogger(cfg.Utils().Log())
		defer logger.SyncLogger()

		servers.NewServer(cfg).Start()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(ComposerCoreCmd)
}
