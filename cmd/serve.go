package cmd

import (
	"github.com/Nattakornn/cache/config"
	"github.com/Nattakornn/cache/modules/servers"
	postgresql "github.com/Nattakornn/cache/pkg/databases/postgressql"
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

		dbPostgresql := postgresql.ConnectDb(cfg.Db())
		defer dbPostgresql.Close()

		servers.NewServer(cfg, dbPostgresql).Start()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(ComposerCoreCmd)
}
